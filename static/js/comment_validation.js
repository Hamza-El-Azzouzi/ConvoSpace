function SubmitComment(event) {
    // 1. Extract Post ID from Current URL
    const pathname = window.location.pathname;
    const postID = pathname.substring(pathname.lastIndexOf('/') + 1)

    // 2. Validation Setup
    const errdisplay = document.getElementById('textarea-error');
    errdisplay.textContent = '';

    // 3. Textarea Validation & Prevent Form Submission if Invalid
    const textareaInput = document.querySelector('textarea[name="textarea"]');
    const comvalue = textareaInput.value.trim();
    if (!comvalue) {
        errdisplay.textContent = 'Comment is required.';
        event.preventDefault();
        return
    }

    // 4. Prevent Default Form Submission
    event.preventDefault();

    // 5. Prepare Form Data
    const formData = textareaInput.value;

    // 6. Send Comment via Fetch API
    fetch("/sendcomment", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ content: formData, postID: postID }),
    })
        .then((response) => {
            // 8. Check Response Validity
            if (!response.ok) {
                throw new Error(`Failed to submit the comment.`);
            }
            return response.json();
        })
        .then((comments) => {
            // 9. Update Comment Section
            UpdateComment(comments);
            textarea.value = "";
        })
        .catch((error) => {
            // 10. Error Handling
            console.error("Error:", error);
        });
}
function UpdateComment(comments) {
    const commentSection = document.querySelector(".comment-section");

    commentSection.innerHTML = "";

    if (comments.length === 0) {
        commentSection.innerHTML = `
          <div class="nothing">
              <p>No comments yet. Be the first to comment!</p>
          </div>`;
        return;
    }

    comments.forEach((comment) => {
        const commentElement = document.createElement("div");
        commentElement.className = "comment";
        commentElement.innerHTML = `
          <div class="comment-header">
              <h6>${comment.Username || "Anonymous"}</h6>
              <i class="fa fa-clock-o">${comment.FormattedDate
            }</i>
          </div>
          <div class="comment-body">${comment.Content}</div>
          <div class="comment-footer">
              <button class="button like" onclick="handleLikeDislike('${comment.CommentID}', 'likeComment', event)">
                  <span id='${comment.CommentID}-likecomment' >👍${comment.LikeCountComment}</span>
              </button>
               <button class="button like" onclick="handleLikeDislike('${comment.CommentID}', 'dislikeComment', event)">
                  <span id='${comment.CommentID}-dislikecomment' >👎${comment.DisLikeCountComment
            }</span>
              </button>
          </div>
      `;
        commentSection.appendChild(commentElement);
    });
}
