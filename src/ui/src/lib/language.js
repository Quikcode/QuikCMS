let nameLocalStorage = "locale"

function Set(value) {
    localStorage.setItem(nameLocalStorage, value)
}

function Get() {
    return localStorage.getItem(nameLocalStorage)
}

export default {
    Set, Get
}