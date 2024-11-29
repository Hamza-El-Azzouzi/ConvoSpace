function SubmitComment(event) {
    const pathname = window.location.pathname;
    const postID = pathname.split('/').pop();
    let isValid = true;

    document.getElementById('textarea-error').textContent = '';

    const textareaInput = document.querySelector('textarea[name="textarea"]');
    if (!textareaInput.value.trim()) {
        document.getElementById('textarea-error').textContent = 'Comment is required.';
        isValid = false;
    }
    
    if (!isValid) {
        event.preventDefault();
        return
    }
    event.preventDefault();
    const formData = textareaInput.value;
      
    fetch("/sendcomment", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ content: formData, postID: postID }),
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error(`Failed to submit the comment.`);
        }
        return response.json();
      })
      .then((comments) => {
        UpdateComment(comments);
        textarea.value = "";
      })
      .catch((error) => {
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
              <i class="fa fa-clock-o">${
                comment.FormattedDate
              }</i>
          </div>
          <div class="comment-body"><pre>${comment.Content}</pre></div>
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
