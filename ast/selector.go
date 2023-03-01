package ast

import (
	"bytes"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
	"strings"
)

type SelectorTerm struct {
	*token.Token
	factor Expression // The identifier of the variable
	fields []string   // The field of the struct (if it is a struct)
	ty     types.Type // The type of the selector term
}

// NewSelectorTerm node
func NewSelectorTerm(factor Expression, fields []string, ty types.Type, tok *token.Token) *SelectorTerm {
	return &SelectorTerm{tok, factor, fields, ty}
}

// String representation of the SelectorTerm node
func (s *SelectorTerm) String() string {
	var out bytes.Buffer

	out.WriteString(s.factor.String())
	// Its possible that the SelectorTerm is an struct field
	for _, field := range s.fields {
		out.WriteString(".")
		out.WriteString(field)
	}

	return out.String()
}

// Build the symbol table for the SelectorTerm node
func (s *SelectorTerm) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Nothing to do here since a SelectorTerm is not added to the symbol table since it is already declared
	return errors
}

// Type check the SelectorTerm node
func (s *SelectorTerm) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	// Type check the factor
	errors = s.factor.TypeCheck(errors, tables, funcEntry)
	// Check if the identifer is declared
	identifer := s.factor.String()
	// Check in the local variables (moves to global scope if not found in local scope)
	entry := funcEntry.Variables.Contains(identifer)

	// If it is not a variable in the local or global scope, check the parameters
	if entry == nil {
		// Check the parameters
		for _, param := range funcEntry.Parameters {
			if param.Name == identifer {
				entry = param
				break
			}
		}
	}
	// If the identifier is not a parameter, local or global variable, then it is an invocation, literal or struct
	if entry != nil {
		// The identifier is a variable/parameter
		// Could be a struct
		if len(s.fields) == 0 {
			// The lvalue is a variable
			s.ty = entry.GetType()
		} else {
			// The lvalue is a struct field
			entryType := entry.GetType()
			var c int
			for i, field := range s.fields {
				c = i
				if types.TypeToKind(entryType) == types.STRUCT {
					entryStructName := entryType.String()[1:] // Remove the * from the type name
					// Check if the struct exists in the symbol table
					structEntry := tables.Structs.Contains(entryStructName)
					if structEntry == nil {
						errors = append(errors, NewSemanticAnalysisError("type "+entryType.String()+" not declared.", "undeclared type", s.Token))
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
							errors = append(errors, NewSemanticAnalysisError("field "+field+" not declared in type "+entryType.String(), "undeclared field", s.Token))

							// Set the type to nil
							s.ty = types.StringToType("nil")
							break
						}

						// The field exists in the struct so move on to the next field
						entryType = matchedField.GetType()
					}
				} else {
					errors = append(errors, NewSemanticAnalysisError("cannot access field "+field+" of non-struct type "+entryType.String(), "invalid field access", s.Token))
					break
				}
			}
			// Primitive type meaning it must be the last field
			if c == len(s.fields)-1 {
				if types.TypeToKind(entryType) == types.STRUCT {
					s.ty = entryType
				} else {
					// Set the type of the lvalue
					s.ty = entryType
				}
			} else {
				s.ty = types.StringToType("nil")
			}
		}
	} else {
		// The identifier is not a variable/parameter
		// The identifier is either an invocation or literal or allocation

		// if identifer is literal
		// The identifier is a literal
		// This is the type of literal
		// s.factor.GetType() returns this

		// if identifer is invocation
		// The identifier is a function invocation
		// Get the return type of the function
		// s.factor.GetType() returns this

		// if identifer is allocation
		// The identifier is an allocation
		// This is a pointer to the struct type being allocated
		// s.factor.GetType() returns this
		s.ty = s.factor.GetType()
	}

	return errors
}

// Get the type of the SelectorTerm node
func (s *SelectorTerm) GetType() types.Type {
	return s.ty
}

// Translate the allocate node into LLVM IR
func (s *SelectorTerm) ToLLVMCFG(tables *st.SymbolTables, blocks []*llvm.BasicBlock, funcEntry *st.FuncEntry, constDecls []*llvm.ConstantDecl) ([]*llvm.BasicBlock, []*llvm.ConstantDecl) {
	// This term can either be a local variable, global variable, parameter, function invocation, literal, allocation or struct field
	// Build the LLVM IR for the factor
	if len(s.fields) == 0 {
		blocks, constDecls = s.factor.ToLLVMCFG(tables, blocks, funcEntry, constDecls)
		// This loads the value of the factor into a register
		// This covers the case of local variable, global variable, parameter, function invocation, literal and allocation
	} else {
		// Only if it is a struct field will the value not be loaded into a register in the above way

		// Get the variable entry
		functionEntry := tables.Funcs.Contains(funcEntry.Name)
		// Check in the local variables (moves to global scope if not found in local scope)
		entry := functionEntry.Variables.Contains(s.factor.String())
		var varName string
		var varType string
		if entry == nil {
			// Check the parameters
			for _, param := range functionEntry.Parameters {
				if param.Name == s.factor.String() {
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
			if field.Name == s.fields[0] {
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

		// Load the value of the field into a register
		loadInst = llvm.NewLoad(llvm.GetPreviousRegister(), fieldEntry.LlvmType)
		loadInst.SetLabel(blocks[len(blocks)-1].GetLabel())
		blocks[len(blocks)-1].AddInstruction(loadInst)

		// Get the type of the field
		fieldType := fieldEntry.GetType().String()[1:] // Remove the * from the type name
		structEntry = tables.Structs.Contains(fieldType)
		varType = fieldEntry.LlvmType // Update the type of the struct

		if strings.Contains(varType, "struct.") {
			varType += "*"
		}

		// Do this for all the remaining fields
		for i := 1; i < len(s.fields); i++ {
			// If there are more fields then they are structs
			// Get the struct field index
			var index int
			for j, field := range structEntry.Fields {
				if field.Name == s.fields[i] {
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

			if i != len(s.fields)-1 {
				// Load the value of the field into a register
				loadInst = llvm.NewLoad(llvm.GetPreviousRegister(), fieldEntry.LlvmType)
				loadInst.SetLabel(blocks[len(blocks)-1].GetLabel())
				blocks[len(blocks)-1].AddInstruction(loadInst)

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
