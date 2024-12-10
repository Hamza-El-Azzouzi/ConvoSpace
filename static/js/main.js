function getCurrentYearAsMax() {
    const currentYear = new Date().getFullYear();
    const  footer = document.querySelector(".footer-part").childNodes[1]
    footer.innerHTML = ` <p>&copy; ${currentYear} Forum01. All Rights Reserved.</p>` 
}
getCurrentYearAsMax()

const clickToLogin = () => {
    window.location.replace('/login');
}