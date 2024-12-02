function validateForm(event) {
  event.preventDefault();

  let email = document.getElementById("email").value.trim();
  let password = document.getElementById("password").value.trim();
  let emailError = document.querySelector(".errorMsg:nth-of-type(1)");
  let passwordError = document.querySelector(".errorMsg:nth-of-type(2)");

  emailError.textContent = "";
  passwordError.textContent = "";

  let isValid = true;

  const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  if (!email) {
    emailError.textContent = "Email is required.";
    isValid = false;
  } else if (!emailPattern.test(email)) {
    emailError.textContent = "Please enter a valid email.";
    isValid = false;
  }

  if (!password) {
    passwordError.textContent = "Password is required.";
    isValid = false;
  }

  if (isValid) {
    fetch("/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email: email, password: password }),
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error(`Failed to submit the login.`);
        }
        return response.json();
      })
      .then((err) => {
        if (err) {
          if (err.activeSession){
            alert(err.activeSession)
            window.location.href = "/";
          }
          if (err.email){
            emailError.textContent = err.email
          }
          if (err.password){
            passwordError.textContent = err.password;
          }
          
        } else {
          window.location.href = "/";
        }
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  }
}

function validateRegistrationForm(event) {
  event.preventDefault();

  const fields = {
    username: {
      value: document.querySelector('input[name="Username"]').value.trim(),
      error: ".errorMsg:nth-of-type(1)",
    },
    email: {
      value: document.querySelector('input[name="email"]').value.trim(),
      error: ".errorMsg:nth-of-type(2)",
    },
    password: {
      value: document.querySelector('input[name="password"]').value.trim(),
      error: ".errorMsg:nth-of-type(3)",
    },
    confirmPassword: {
      value: document
        .querySelector('input[name="Confirm-Password"]')
        .value.trim(),
      error: ".errorMsg:nth-of-type(4)",
    },
  };

  const passwordRegex =
    /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[^a-zA-Z\d]).{7,}$/;

  document
    .querySelectorAll(".errorMsg")
    .forEach((msg) => (msg.textContent = ""));

  let isValid = true;

  if (!fields.username.value) {
    showError(fields.username.error, "Username is required.");
    isValid = false;
  }

  const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  if (!fields.email.value) {
    showError(fields.email.error, "Email is required.");
    isValid = false;
  } else if (!emailPattern.test(fields.email.value)) {
    showError(fields.email.error, "Invalid email! ");
    isValid = false;
  }

  if (!fields.password.value) {
    showError(fields.password.error, "Password is required.");
    isValid = false;
  } else if (!passwordRegex.test(fields.password.value)) {
    showError(
      fields.password.error,
      `-Password must be at least 8 characters long.<br>-include one uppercase letter.<br>-one lowercase letter.<br>-one number.<br>-and one special character.`
    );
    isValid = false;
  }

  if (!fields.confirmPassword.value) {
    showError(fields.confirmPassword.error, "Confirm password is required.");
    isValid = false;
  } else if (fields.password.value !== fields.confirmPassword.value) {
    showError(fields.confirmPassword.error, "Passwords do not match!");
    isValid = false;
  }

  if (isValid) {
    fetch("/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ username : fields.username.value ,email: fields.email.value, password: fields.password.value , confirmPassword : fields.confirmPassword.value }),
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error(`Failed to submit the register.`);
        }
        return response.json();
      })
      .then((err) => {
        if (err) {
          if (err.activeSession){
            alert(err.activeSession)
            window.location.href = "/";
          }
          if (err.username){
            const errPart =  document.querySelector(fields.username.error)
            showError(fields.username.error,err.username)
          }
          if (err.email){
            const errPart =  document.querySelector(fields.email.error)
            console.log(errPart)
            showError(fields.email.error,err.email)
          }
          if (err.password){
            const errPart =  document.querySelector(fields.password.error)
            console.log(errPart)
            showError(fields.password.error,err.password)
          }
          if (err.conPassowrd){
            const errPart =  document.querySelector(fields.confirmPassword.error)
            console.log(errPart)
            showError(fields.confirmPassword.error,err.confirmPassword)
          }
        } else {
          window.location.href = "/";
        }

      })
      .catch((error) => {
        console.error("Error:", error);
      });
  }
}

function showError(selector, message) {
  document.querySelector(selector).innerHTML = message.replace(/\n/g, "<br>");
}
