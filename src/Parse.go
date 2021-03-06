package src

import (
    "strings"
    "strconv"
    //"fmt"
)

var Variables = make(map[string]Variable);
var Reserved = []string{"clear", "clean", "exit"};

func Transact(query string) string {
    return eval(query);
}

// Evaluates user input to return
func eval(query string) string {
    query = strings.Replace(query, " ", "", -1);
    query = replaceSubtraction(query);
    if (checkPrimitive(query)) {
        return query;
    }

    if (IsVariable(query)) {
        val, contains := Variables[query];
        if (!contains) {
            return "ERROR: Undefined variable: '" + query + "'";
        }
        return Get(val);
    }

    if (isFunctionCall(query)) {
        funcNameIndex := strings.Index(query, "(");
        funcName := query[0:funcNameIndex+1];
        _, match := Functions[funcName]
        if (match) {
            if (!strings.Contains(query, "=")) {
                matrix, err := ApplyFunction(query);
                if (err != "") {
                    return err;
                }
                return Print(matrix);
            } else {
                return assignVariable(query);
            }
        }
        _, match = IntFunctions[funcName]
        if (match) {
            if (!strings.Contains(query, "=")) {
                val, err := ApplyIntFunction(query);
                if (err != "") {
                    return err;
                }
                return strconv.Itoa(val);
            } else {
                return assignVariable(query);
            }
        }
        _, match = FloatFunctions[funcName];
        if (match) {
            if (!strings.Contains(query, "=")) {
                val, err := ApplyFloatFunction(query);
                if (err != "") {
                    return err;
                }
                return strconv.FormatFloat(val, 'f', -1, 64);
            } else {
                return assignVariable(query);
            }
        }
    }

    if (containsMatrix(query)) {
        if (string(query[0]) == "[") {
            rows := queryToValues(query);
            matrix, err := NewMatrix(rows);
            if (err != "") {
                return err;
            } else {
                return Print(matrix);
            }
        }

        if (strings.Contains(query, "=")) {
            return assignVariable(query);
        }

        numArgs := parseArithmetic(query);

        if (numArgs == 1) {
            _, str := ApplyMatrixOperation(query);
            return str;
        } else {
            _, str := ApplyMultipleMatrixOperations(query);
            return str;
        }

    } else if (strings.Contains(query, "=")) {
        return assignVariable(query);
    } else if (parseArithmetic(query) == 1) {
        return ApplyArithmetic(query);
    } else if (parseArithmetic(query) > 1) {
        return ApplyMultipleArithmetic(query);
    } else {
        return "ERROR: Malformed query";
    }
}

// Tests if query contains an arithmetic argument
func parseArithmetic(query string) int {
    matched := CountAny(query, "+","~", "*", "/");
    return matched;
}

// Currently, primitives are defined as integers and floats
func checkPrimitive(query string) bool {
    query = strings.TrimSpace(query);
    _, err := strconv.Atoi(query);
    if (err != nil) {
        _, err2 := strconv.ParseFloat(query, 64)
            if (err2 != nil) {
                return false;
            }
            return true;
        }
    return true;
}

func CountAny(str string, seps ...string) (i int) {
    for _, sep := range(seps) {
        i += strings.Count(str, sep);
    }
    return i;
}

// Replaces "-" with "~" as the subtraction operator so as not to confuse subtraction with a negative sign
func replaceSubtraction(query string) string {
    firstInst := strings.Index(query, "-");
    if (firstInst == -1) {
        return query;
    }
    if (firstInst != 0 && !strings.ContainsAny(string(query[firstInst - 1]), "+ & * & ~ & / & [ & ; & ,")) {
        query = strings.Replace(query, "-", "~", 1);
        return replaceSubtraction(query);
    } else if (firstInst == 0) { // must be a number
        replaced := []string{"-", replaceSubtraction(query[1:len(query)])};
        return strings.Join(replaced, "");
    } else if (strings.ContainsAny(string(query[firstInst - 1]), "+ & * & ~ & /")) {
        // If prev byte is an operator, current byte must be a number
        replaced := []string{string(query[0]), replaceSubtraction(query[1:len(query)])};
        return strings.Join(replaced, "");
    } else if (string(query[firstInst-1]) == "(") {
        return query;
    } else {
        return query;
    }
}
