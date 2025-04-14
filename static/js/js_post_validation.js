

document.getElementById('postForm').addEventListener('submit', function (event) {
    let isValid = true;

    document.getElementById('title-error').textContent = '';
    document.getElementById('category-error').textContent = '';
    document.getElementById('textarea-error').textContent = '';
    document.getElementById('image-error').textContent = '';

    const titleInput = document.querySelector('input[name="title"]');
    if (titleInput.value.length > 250) {
        document.getElementById('title-error').textContent = 'Maximum 255 characters.';
        isValid = false;
    }
    if (!titleInput.value.trim()) {
        document.getElementById('title-error').textContent = 'Title is required.';
        isValid = false;
    }

    const categoryCheckboxes = document.querySelectorAll('input[name="category"]:checked');
    if (categoryCheckboxes.length === 0) {
        document.getElementById('category-error').textContent = 'Please select at least one category.';
        isValid = false;
    }

    const textareaInput = document.querySelector('textarea[name="textarea"]');
    if (textareaInput.value.length > 10000) {
        document.getElementById('textarea-error').textContent = "Maximum 10000 characters.";
        isValid = false;
    }
    if (!textareaInput.value.trim()) {
        document.getElementById('textarea-error').textContent = 'Description is required.';
        isValid = false;
    }
    const imageInput = document.getElementById("imageUpload");
    const maxSize = 20 * 1024 * 1024;
    const allowedTypes = ['image/jpeg', 'image/png', 'image/gif'];
    const file = imageInput.files[0];
    if (file) {
        if (file.size > maxSize) {
            document.getElementById('image-error').textContent = 'File size must be less than 0.5MB.';
            isValid = false;
        }
        if (!allowedTypes.includes(file.type)) {
            document.getElementById('image-error').textContent = 'Only JPG, PNG and GIF files are allowed.';
            isValid = false;
        }
    }


    if (!isValid) {
        event.preventDefault();
    }
});