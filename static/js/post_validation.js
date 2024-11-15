

document.getElementById('postForm').addEventListener('submit', function (event) {
    let isValid = true;

    document.getElementById('title-error').textContent = '';
    document.getElementById('category-error').textContent = '';
    document.getElementById('textarea-error').textContent = '';


    const titleInput = document.querySelector('input[name="title"]');
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
    if (!textareaInput.value.trim()) {
        document.getElementById('textarea-error').textContent = 'Description is required.';
        isValid = false;
    }

    if (!isValid) {
        event.preventDefault();
    }
});