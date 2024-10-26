package logger

func contains(arr []string, s string) bool {
    for _, item := range arr {
        if item == s {
            return true
        }
    }

    return false
}

