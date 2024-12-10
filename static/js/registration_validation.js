const username = document.querySelector("input[name='username']")
const email = document.querySelector("input[name='email']")
const password = document.querySelector("input[name='passwd']")
const confirmPassword = document.querySelector("input[name='confirmPasswd']")
const button = document.querySelector("input[type='submit']")

const ErrMessageName = document.getElementById("nameErr")
const ErrMessageEmail = document.getElementById("emailErr")
const ErrMessagePasswd = document.getElementById("passwdErr")
const ErrMessageConfirmPasswd = document.getElementById('confirmPasswdErr')
const Err = document.getElementById("otherErr")

const ExpName = /^[a-zA-Z0-9_]{3,20}$/
const ExpEmail = /^[a-zA-Z0-9._+-=]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/
const ExpPasswd = /^(?=(.*[a-z]))(?=(.*[A-Z]))(?=(.*[0-9]))(?=(.*[^a-zA-Z0-9]))(.{7,})$/


const InvalidEmail = "invalid email!! enter a valid email"
const InvalidName = "invalid name!! enter a valid name"
const InvalidPsswd = "invalid password!! enter a valid password"
const NotMatch = "password confirmation doesn't match!!"

const Error = (elem, errorMssg) => {
    elem.textContent = errorMssg
    elem.style.color = "red"
    elem.style.fontSize = "12px"
}
// const ResetErrorTags = () => {

// }
const VerifyData = () => {
    let exist = false
    ErrMessageName.textContent = ""
    ErrMessageEmail.textContent = ""
    ErrMessagePasswd.textContent = ""
    ErrMessageConfirmPasswd.textContent = ""
    Err.textContent = ""
    switch (true) {
        case (!ExpName.test(username.value)):
            console.log("test ", ExpName.test(username.value), username.value, username.value.length)
            Error(ErrMessageName, InvalidName)
            exist = true
            break
        case (!ExpEmail.test(email.value)):
            Error(ErrMessageEmail, InvalidEmail)
            exist = true
            break
        case (!ExpPasswd.test(password.value)) || (!ExpPasswd.test(confirmPassword.value)):
            Error(ErrMessagePasswd, InvalidPsswd)
            exist = true
            break
        case (password.value !== confirmPassword.value):
            console.log(password.value, confirmPassword.value)
            Error(ErrMessageConfirmPasswd, NotMatch)
            exist = true
    }
    return exist
}

button.addEventListener("click", (event) => {
    event.preventDefault()
    console.log("verify ", VerifyData())
    if (!VerifyData()) {
        Err.textContent = ""
        console.log(username.value, email.value, password.value, confirmPassword.value)

        fetch("/register", {
            headers: {
                "Content-Type": "application/json",
            },
            method: "POST",
            body: JSON.stringify({
                username: username.value,
                email: email.value,
                password: password.value,
                confirmPassword: confirmPassword.value
            }),
        })
            .then(response => response.json())
            .then(reply => {
                switch (true) {
                    case (reply.REplyMssg == "Sign Done"):
                        window.location.href = "/login"
                        break
                    case (reply.REplyMssg == "err"):
                        Error(Err, "Username or Email already exist!!")
                        break
                    case (reply.REplyMssg == "notMatch"):
                        Error(ErrMessageConfirmPasswd, NotMatch)
                }
            })
    }


})