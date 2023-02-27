package ast

import (
	"bytes"
	"fmt"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
	"strings"
)

// LValue node in the AST
type LValue struct {
	*token.Token
	identifier string     // The identifier of the variable
	fields     []string   // The field of the struct
	ty         types.Type // The type of the lvalue
}

// NewLValue node
func NewLValue(identifier string, fields []string, ty types.Type, tok *token.Token) *LValue {
	return &LValue{tok, identifier, fields, ty}
}

// String representation of the lvalue node
func (l *LValue) String() string {
	var out bytes.Buffer

	out.WriteString(l.identifier)
	// Its possible that the lvalue is an struct field
	for _, field := range l.fields {
		out.WriteString(".")
		out.WriteString(field)
	}

	return out.String()
}

// Build the symbol table for the lvalue node
func (l *LValue) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Nothing to do here since a lvalue is not added to the symbol table since it is already declared
	return errors
}

// Type check the lvalue node
func (l *LValue) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	// Check if the identifer is declared
	entry := funcEntry.Variables.Contains(l.identifier)
	if entry == nil {
		// Check the parameters
		for _, param := range funcEntry.Parameters {
			if param.Name == l.identifier {
				entry = param
				break
			}
		}
	}
	if entry == nil {
		errors = append(errors, NewSemanticAnalysisError("undeclared variable "+l.identifier+".", "undeclared variable", l.Token))
	} else {
		if len(l.fields) == 0 {
			// The lvalue is a variable
			l.ty = entry.GetType()
		} else {
			entryType := entry.GetType()
			var c int
			for i, field := range l.fields {
				c = i
				if types.TypeToKind(entryType) == types.STRUCT {
					entryStructName := entryType.String()[1:] // Remove the * from the type name
					// Check if the struct exists in the symbol table
					structEntry := tables.Structs.Contains(entryStructName)
					if structEntry == nil {
						errors = append(errors, NewSemanticAnalysisError("type "+entryType.String()+" not declared.", "undeclared type", l.Token))
						break
					} else {
						// Check if the field exists in the struct
						found := false

						var matchedField *st.FieldEntry
						for _, fieldEntry := range structEntry.Fields {
							if fieldEntry.Name == field {
								found = true
								matchedField = fieldEntry
								break
							}
						}
						if !found {
							errors = append(errors, NewSemanticAnalysisError("field "+field+" not declared in type "+entryType.String(), "undeclared field", l.Token))

							// Set the type to nil
							l.ty = types.StringToType("nil")
							break
						}

						// The field exists in the struct so move on to the next field
						entryType = matchedField.GetType()
					}
				} else {
					fmt.Println(l.identifier, l.fields)
					errors = append(errors, NewSemanticAnalysisError("cannot access field "+field+" of non-struct type "+entryType.String(), "invalid field access", l.Token))
					break
				}
			}
			// Primitive type meaning it must be the last field
			if c == len(l.fields)-1 {
				if types.TypeToKind(entryType) == types.STRUCT {
					l.ty = entryType
				} else {
					// Set the type of the lvalue
					l.ty = entryType
				}
			} else {
				l.ty = types.StringToType("nil")
			}
		}
	}

	return errors
}

// Get the type of the lvalue node
func (l *LValue) GetType() types.Type {
	return l.ty
}

// Translate the lvalue node into LLVM IR
func (l *LValue) ToLLVMCFG(tables *st.SymbolTables, blocks []*llvm.BasicBlock, funcEntry *st.FuncEntry, constDecls []*llvm.ConstantDecl) ([]*llvm.BasicBlock, []*llvm.ConstantDecl) {
	// Stay in the same block

	// The lvalue could either be a variable or a field of a struct
	// Check if the lvalue is a variable
	if len(l.fields) == 0 {
		// The lvalue is a variable
		// Get the variable entry
		functionEntry := tables.Funcs.Contains(funcEntry.Name)
		entry := functionEntry.Variables.Contains(l.identifier)
		if entry == nil {
			for _, param := range funcEntry.Parameters {
				if param.Name == l.identifier {
					entry = param
					break
				}
			}

			// Load the parameter into a register
			// Make the load instruction
			ty := entry.LlvmTy
			if strings.Contains(ty, "struct.") {
				ty += "*"
			}
			loadInst := llvm.NewLoad("%P_"+entry.Name, ty)
			loadInst.SetLabel(blocks[len(blocks)-1].GetLabel())
			blocks[len(blocks)-1].AddInstruction(loadInst)

		} else {
			// Load the variable into a register
			var varName string
			if entry.Scope == st.GLOBAL {
				varName = "@" + entry.Name
			} else {
				varName = "%" + entry.Name
			}

			ty := entry.LlvmTy
			if strings.Contains(ty, "struct.") {
				ty += "*"
			}

			loadInst := llvm.NewLoad(varName, ty)
			loadInst.SetLabel(blocks[len(blocks)-1].GetLabel())
			blocks[len(blocks)-1].AddInstruction(loadInst)
		}
	} else {
		// The lvalue is a field of a struct
		// Get the variable entry
		functionEntry := tables.Funcs.Contains(funcEntry.Name)
		// Check in the local variables (moves to global scope if not found in local scope)
		entry := functionEntry.Variables.Contains(l.identifier)
		var varName string
		var varType string
		if entry == nil {
			// Check the parameters
			for _, param := range functionEntry.Parameters {
				if param.Name == l.identifier {
					entry = param
					varName = "%P_" + param.Name
					varType = entry.LlvmTy
					break
				}
			}
		} else {
			if entry.Scope == st.GLOBAL {
				varName = "@" + entry.Name
			} else {
				varName = "%" + entry.Name
			}
			varType = entry.LlvmTy
		}

		if strings.Contains(varType, "struct.") {
			varType += "*"
		}

		// Load the variable into a register
		loadInst := llvm.NewLoad(varName, varType)
		loadInst.SetLabel(blocks[len(blocks)-1].GetLabel())
		blocks[len(blocks)-1].AddInstruction(loadInst)

		// Get the struct
		entryStructName := entry.GetType().String()[1:] // Remove the * from the type name
		structEntry := tables.Structs.Contains(entryStructName)

		// Get the struct field index
		var index int
		var fieldEntry *st.FieldEntry
		for i, field := range structEntry.Fields {
			if field.Name == l.fields[0] {
				index = i
				fieldEntry = field
				break
			}
		}

		// Add the getelementptr instruction
		// Make sure to indicate that the struct is a pointer (in varType)
		selectElementInst := llvm.NewGetElementPtr(llvm.GetPreviousRegister(), varType, index)
		selectElementInst.SetLabel(blocks[len(blocks)-1].GetLabel())
		blocks[len(blocks)-1].AddInstruction(selectElementInst)

		// Get the type of the field
		fieldType := fieldEntry.GetType().String()[1:] // Remove the * from the type name
		structEntry = tables.Structs.Contains(fieldType)
		varType = fieldEntry.LlvmType // Update the type of the struct

		if strings.Contains(varType, "struct.") {
			varType += "*"
		}

		// Do this for all the remaining fields
		for i := 1; i < len(l.fields); i++ {
			// If there are more fields then they are structs
			// Get the struct field index
			var index int
			for j, field := range structEntry.Fields {
				if field.Name == l.fields[i] {
					index = j
					fieldEntry = field
					break
				}
			}

			// Add the getelementptr instruction
			// Make sure to indicate that the struct is a pointer (in varType)
			selectElementInst = llvm.NewGetElementPtr(llvm.GetPreviousRegister(), varType, index)
			selectElementInst.SetLabel(blocks[len(blocks)-1].GetLabel())
			blocks[len(blocks)-1].AddInstruction(selectElementInst)

			if i != len(l.fields)-1 {
				// Get the type of the field
				fieldType = fieldEntry.GetType().String()[1:] // Remove the * from the type name
				structEntry = tables.Structs.Contains(fieldType)
				varType = fieldEntry.LlvmType // Update the type of the struct

				if strings.Contains(varType, "struct.") {
					varType += "*"
				}
			}
		}
	}
	return blocks, constDecls
}
