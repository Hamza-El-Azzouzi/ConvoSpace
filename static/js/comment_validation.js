document.getElementById('commentForm').addEventListener('submit', function (event) {
    let isValid = true;

    document.getElementById('textarea-error').textContent = '';

    const textareaInput = document.querySelector('textarea[name="textarea"]');
    if (!textareaInput.value.trim()) {
        document.getElementById('textarea-error').textContent = 'Comment is required.';
        isValid = false;
    }
    
    if (!isValid) {
        event.preventDefault();
    }
});