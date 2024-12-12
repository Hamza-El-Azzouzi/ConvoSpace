function getCurrentYearAsMax() {
    const currentYear = new Date().getFullYear();
    const  footer = document.querySelector(".footer-part").childNodes[1]
    footer.innerHTML = ` <p>&copy; ${currentYear} 01-Forum. All Rights Reserved.</p>` 
}
getCurrentYearAsMax()

const clickToLogin = () => {
    window.location.replace('/login');
}