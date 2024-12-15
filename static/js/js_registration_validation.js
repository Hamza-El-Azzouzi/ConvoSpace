const username = document.querySelector("input[name='username']")
const email = document.querySelector("input[name='email']")
const password = document.querySelector("input[name='passwd']")
const confirmPassword = document.querySelector("input[name='confirmPasswd']")
const button = document.querySelector("input[type='submit']")

const ErrMessageName = document.getElementById("nameErr")
const ErrMessageEmail = document.getElementById("emailErr")
const ErrMessagePasswd1st = document.getElementById("passwdErr1st")
const ErrMessagePasswd = document.getElementById("passwdErr2nd")
const ErrMessageConfirmPasswd = document.getElementById('confirmPasswdErr')
const Err = document.getElementById("otherErr")

const ExpName = /^[a-zA-Z0-9_.]{3,20}$/
const ExpEmail = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/
const ExpPasswd = /^(?=(.*[a-z]))(?=(.*[A-Z]))(?=(.*[0-9]))(?=(.*[^a-zA-Z0-9]))(.{8,20})$/


const InvalidEmail = "invalid email!! enter a valid email"
const InvalidName = "invalid name!! enter a valid name"
const Invalid = "invalid password!!"
const NotMatch = "password confirmation doesn't match!!"

const InvalidPsswd = () => {
    ErrMessagePasswd1st.innerHTML = `<h4 style="text-align: center">--------Password Requirements--------</h4>
    <ul>
        <li>At least one <strong>lowercase letter</strong> (a-z)</li>
        <li>At least one <strong>uppercase letter</strong> (A-Z)</li>
        <li>At least one <strong>digit</strong> (0-9)</li>
        <li>At least one <strong>special character</strong> (anything not a letter or a digit)</li>
        <li>Password must be at least <strong>8 characters long</strong></li>
    </ul>`
    ErrMessagePasswd1st.style.textAlign = "left"
    ErrMessagePasswd1st.style.color = "red"
    ErrMessagePasswd1st.style.marginTop = "0px"
}
const Error = (elem, errorMssg) => {
    elem.textContent = errorMssg
    elem.style.color = "red"
    elem.style.fontSize = "12px"
    elem.style.marginTop = "0px"
}
const VerifyData = () => {
    let exist = false

    ErrMessageName.textContent = ""
    ErrMessageEmail.textContent = ""
    ErrMessagePasswd.textContent = ""
    ErrMessageConfirmPasswd.textContent = ""
    ErrMessagePasswd1st.textContent = ""
    Err.textContent = ""

    switch (true) {
        case (!ExpName.test(username.value)):
            Error(ErrMessageName, InvalidName)
            exist = true
            break
        case (!ExpEmail.test(email.value)):
            Error(ErrMessageEmail, InvalidEmail)
            exist = true
            break
        case (!ExpPasswd.test(password.value)):
            InvalidPsswd("passwdErr1st")
            exist = true
            break
        case (!ExpPasswd.test(confirmPassword.value)):
            Error(ErrMessagePasswd, Invalid)
            exist = true
            break
        case (password.value !== confirmPassword.value):
            Error(ErrMessageConfirmPasswd, NotMatch)
            exist = true
    }
    return exist
}

button.addEventListener("click", (event) => {
    event.preventDefault()

    if (!VerifyData()) {
        Err.textContent = ""

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
                    case (reply.REplyMssg == "session"):
                        Error(Err, "you already have an active session")
                        break
                    case (reply.REplyMssg == "Done"):
                        window.location.href = "/login"
                        break
                    case (reply.REplyMssg == "email"):
                        Error(ErrMessageEmail, "Email already exist!!")
                        break
                    case (reply.REplyMssg == "user"):
                        Error(ErrMessageName, "Username already exist!!")
                        break
                    case (reply.REplyMssg == "notMatch"):
                        Error(ErrMessageConfirmPasswd, NotMatch)
                }
            })
    }


})