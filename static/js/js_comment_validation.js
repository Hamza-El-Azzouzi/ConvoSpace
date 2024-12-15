let commentOffset = 1;
const commentsPerPage = 5;
const btn = document.querySelector('.load-more-btn'); 
function SubmitComment(event) {

    event.preventDefault();

    const path = window.location.pathname;
    const currentPostId = path.substring(path.lastIndexOf('/') + 1);
    const textarea = document.querySelector('.comment-textarea');

    if (textarea.value.trim() === "") {
        document.getElementById('textarea-error').textContent = 'Comment is required.';
        return
    }
    const currentValue = textarea.value

    async function fetchData() {
        try {
            const response = await fetch('/sendcomment', {
                method: "POST",
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ content: currentValue, postId: currentPostId})
            })
            if (!response.ok) {
                throw new Error('http error')
            }
            
            textarea.value = ''
            const update = await response.json()
            UpdateComment(update, false)
        } catch (error) {
            console.error('There was a problem in fetch :', error);
        }
    }
    fetchData()
}

function UpdateComment(comments, append = false) {
    const commentSection = document.querySelector(".comment-section");
    if (!append) {
        commentSection.innerHTML = ""; 
    }

    if (comments.length === 0 && !append) {
        const noCommentsDiv = document.createElement('div');
        noCommentsDiv.className = 'nothing';
        noCommentsDiv.innerHTML = '<p>No comments yet. Be the first to comment!</p>';
        document.querySelector('.load-more-btn').style.display = 'none'; 
        commentSection.appendChild(noCommentsDiv);
        return;
    }

    comments.forEach((comment) => {
        const { Username = "Anonymous", FormattedDate, Content, CommentID, LikeCountComment, DisLikeCountComment } = comment;
        const commentElement = document.createElement("div");
        commentElement.className = "comment";
        commentElement.innerHTML = `
          <div class="comment-header">
              <h6>${Username}</h6>
              <i class="fa fa-clock-o">${FormattedDate}</i>
          </div>
          <div class="comment-body"><pre>${Content}</pre></div>
          <div class="comment-footer">
              <button class="button like" onclick="handleLikeDislike('${CommentID}', 'likeComment', event)">
                  <span id='${CommentID}-likecomment'>👍${LikeCountComment}</span>
              </button>
               <button class="button like" onclick="handleLikeDislike('${CommentID}', 'dislikeComment', event)">
                  <span id='${CommentID}-dislikecomment'>👎${DisLikeCountComment}</span>
              </button>
          </div>
      `;
        commentSection.appendChild(commentElement);
    });
}
async function loadMore() {
    const postID = window.location.href.split("/")[4];
    const queryParams = new URLSearchParams({
        postId: postID,
        offset: commentOffset * commentsPerPage,
    });
    try {
        const response = await fetch(`/comment?${queryParams}`);
        if (!response.ok) {
            throw new Error('Failed to load more comments');
        }

        const newComments = await response.json();

        if (newComments.length < commentsPerPage) {
            btn.style.display = 'none'; 
        }

        UpdateComment(newComments, true);

        commentOffset += 1;
    } catch (error) {
        console.error('Error fetching more comments:', error);
    }
}
if(btn){
    btn.addEventListener('click', loadMore);
}