function validateForm(event) {
    event.preventDefault(); 

    let email = document.getElementById('email').value.trim();
    let password = document.getElementById('password').value.trim();
    let emailError = document.querySelector('.errorMsg:nth-of-type(1)');
    let passwordError = document.querySelector('.errorMsg:nth-of-type(2)');

    emailError.textContent = '';
    passwordError.textContent = '';

    let isValid = true;

    const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!email) {
      emailError.textContent = 'Email is required.';
      isValid = false;
    } else if (!emailPattern.test(email)) {
      emailError.textContent = 'Please enter a valid email.';
      isValid = false;
    }

    if (!password) {
      passwordError.textContent = 'Password is required.';
      isValid = false;
    }

    if (isValid) {
      event.target.submit();
    }
}

function validateRegistrationForm(event) {
  event.preventDefault();

  const fields = {
    username: { value: document.querySelector('input[name="Username"]').value.trim(), error: '.errorMsg:nth-of-type(1)' },
    email: { value: document.querySelector('input[name="email"]').value.trim(), error: '.errorMsg:nth-of-type(2)' },
    password: { value: document.querySelector('input[name="password"]').value.trim(), error: '.errorMsg:nth-of-type(3)' },
    confirmPassword: { value: document.querySelector('input[name="Confirm-Password"]').value.trim(), error: '.errorMsg:nth-of-type(4)' },
  };

  const passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[^a-zA-Z\d]).{7,}$/;

  document.querySelectorAll('.errorMsg').forEach(msg => (msg.textContent = ''));

  let isValid = true;

  if (!fields.username.value) {
    showError(fields.username.error, 'Username is required.');
    isValid = false;
  }

  const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  if (!fields.email.value) {
    showError(fields.email.error, 'Email is required.');
    isValid = false;
  } else if (!emailPattern.test(fields.email.value)) {
    showError(fields.email.error, 'Invalid email! ');
    isValid = false;
  }

  if (!fields.password.value) {
    showError(fields.password.error, 'Password is required.');
    isValid = false;
  } else if (!passwordRegex.test(fields.password.value)) {
    showError(fields.password.error, `-Password must be at least 8 characters long.<br>-include one uppercase letter.<br>-one lowercase letter.<br>-one number.<br>-and one special character.`);
    isValid = false;
  }

  console.log(fields.confirmPassword.value)

  if (!fields.confirmPassword.value) {
    showError(fields.confirmPassword.error, 'Confirm password is required.');
    isValid = false;
  } else if (fields.password.value !== fields.confirmPassword.value) {
    showError(fields.confirmPassword.error, 'Passwords do not match!');
    isValid = false;
  }

  if (isValid) {
    event.target.submit();
  }
}

function showError(selector, message) {
  document.querySelector(selector).innerHTML = message.replace(/\n/g, '<br>');
}