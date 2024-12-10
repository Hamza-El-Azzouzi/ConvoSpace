// function validateForm(event) {

// }

// function validateRegistrationForm(event) {

// }

// function showError(selector, message) {

// }
const username = document.querySelector("input[name='username']")
const email = document.querySelector("input[name='email']")
const password = document.querySelector("input[name='passwd']")
const submitBtn = document.querySelector("input[type='submit']")
const ErrMessageEmail = document.getElementById("emailErr")
const ErrMessagePasswd = document.getElementById("passwdErr")
// const Err = document.getElementById("otherErr")
const InvalidEmail = "invalid email!! enter a valid email"
const InvalidPsswd = "invalid password!! enter a valid password"
const ExpEmail = /^[a-zA-Z0-9._+-=]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/
const ExpPasswd = /^(?=(.*[a-z]))(?=(.*[A-Z]))(?=(.*[0-9]))(?=(.*[^a-zA-Z0-9]))(.{7,})$/

const Error = (elem, errorMssg) => {
    elem.textContent = errorMssg
    elem.style.color = "red"
    elem.style.fontSize = "12px"
}

const VerifyData = () => {
    ErrMessageEmail.textContent = ""
    ErrMessagePasswd.textContent = ""
    // Err.textContent = ""
    console.log("PASSWD", ExpPasswd.test(password.value))
    console.log("value ", password.value, "value2", email.value)
    switch (true) {
        case (!ExpEmail.test(email.value)):
            Error(ErrMessageEmail, InvalidEmail)
            break
        case (!ExpPasswd.test(password.value)):
            Error(ErrMessagePasswd, InvalidPsswd)
    }
}

submitBtn.addEventListener("click", (event) => {
    event.preventDefault()
    VerifyData()

    fetch("/login", {
        headers: {
            "Content-Type": "application/json",
        },
        method: "POST",
        body: JSON.stringify({ email: email.value, password: password.value, })
    })
        .then(response => response.json())
        .then(reply => {
            switch (true) {
                case (reply.REplyMssg == "login Done"):
                    window.location.href = "/"
                    break
                case (reply.REplyMssg == "email"):
                    Error(ErrMessageEmail, "email not found!!, create an account")
                    break
                case (reply.REplyMssg == "passdw"):
                    Error(ErrMessagePasswd, "incorrect Password!!, TRy again")
                // case (reply.REplyMssg == "session"):
                //    Error(Err, "you already have an active session")
            }
        })
})