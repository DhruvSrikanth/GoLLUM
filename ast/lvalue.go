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
func (l *LValue) ToLLVMCFG(tables *st.SymbolTables, blocks []*llvm.BasicBlock, funcEntry *st.FuncEntry, constDecls []*llvm.ConstantDecl) ([]*llvm.BasicBlock, []*llvm.ConstantDecl, string) {
	// Stay in the same block
	var mostRecentOperand string

	// The lvalue could either be a variable or a field of a struct
	if len(l.fields) == 0 {
		// The lvalue is a variable
		if entry := funcEntry.Variables.Contains(l.identifier); entry != nil {
			// lvalue is a local or global variable
			if entry.Scope == st.GLOBAL {
				// Global variable
				mostRecentOperand = "@" + entry.Name
			} else {
				// Local variable
				mostRecentOperand = "%" + entry.Name
			}
		} else {
			// lvalue is a parameter
			for _, param := range funcEntry.Parameters {
				if param.Name == l.identifier {
					mostRecentOperand = "%P_" + param.Name
					break
				}
			}
		}
	} else {
		// The lvalue is a field of a struct
		var varName string
		var varType string
		var varEntry *st.VarEntry
		// Check if it is a parameter or variable (local and global)
		if varEntry = funcEntry.Variables.Contains(l.identifier); varEntry != nil {
			if varEntry.Scope == st.GLOBAL {
				varName = "@" + varEntry.Name
			} else {
				varName = "%" + varEntry.Name
			}
			varType = varEntry.LlvmTy
		} else {
			// Check the parameters
			for _, param := range funcEntry.Parameters {
				if param.Name == l.identifier {
					varEntry = param
					varName = "%P_" + param.Name
					varType = varEntry.LlvmTy
					break
				}
			}
		}

		// Load the variable into a register
		if strings.Contains(varType, "struct.") {
			varType += "*"
		}
		loadInst := llvm.NewLoad(varName, varType)
		loadInst.SetLabel(blocks[len(blocks)-1].GetLabel())
		blocks[len(blocks)-1].AddInstruction(loadInst)
		mostRecentOperand = llvm.GetPreviousRegister()

		// Get the struct entry in the symbol table
		entryStructName := varEntry.GetType().String()[1:] // Remove the * from the type name
		structEntry := tables.Structs.Contains(entryStructName)

		// Get the struct field index to use it as the offset
		var fieldEntry *st.FieldEntry
		var fieldIndex int

		fieldEntry, fieldIndex = structEntry.GetField(l.fields[0])

		// Add the getelementptr instruction
		if strings.Contains(varType, "struct.") {
			// Remove the * from the type name
			varType = varType[:len(varType)-1]
		}
		selectElementInst := llvm.NewGetElementPtr(mostRecentOperand, varType, fieldIndex)
		selectElementInst.SetLabel(blocks[len(blocks)-1].GetLabel())
		blocks[len(blocks)-1].AddInstruction(selectElementInst)
		mostRecentOperand = llvm.GetPreviousRegister()

		// Do this for all the remaining fields
		for i := 1; i < len(l.fields); i++ {
			// Get the type of the field
			fieldType := fieldEntry.GetType().String()[1:] // Remove the * from the type name
			structEntry = tables.Structs.Contains(fieldType)
			varType = fieldEntry.LlvmType // Update the type of the struct

			// First load the struct in to a register to get a pointer to it
			if strings.Contains(varType, "struct.") {
				varType += "*"
			}
			loadInst := llvm.NewLoad(mostRecentOperand, varType)
			loadInst.SetLabel(blocks[len(blocks)-1].GetLabel())
			blocks[len(blocks)-1].AddInstruction(loadInst)
			mostRecentOperand = llvm.GetPreviousRegister()

			// Get the struct field index
			fieldEntry, fieldIndex = structEntry.GetField(l.fields[i])

			// Add the getelementptr instruction
			if strings.Contains(varType, "struct.") {
				// Remove the * from the type name
				varType = varType[:len(varType)-1]
			}
			selectElementInst = llvm.NewGetElementPtr(mostRecentOperand, varType, fieldIndex)
			selectElementInst.SetLabel(blocks[len(blocks)-1].GetLabel())
			blocks[len(blocks)-1].AddInstruction(selectElementInst)
			mostRecentOperand = llvm.GetPreviousRegister()
		}
	}
	return blocks, constDecls, mostRecentOperand
}
