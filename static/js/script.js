document.addEventListener("DOMContentLoaded", function () {
    fetch("/Posts")
      .then((response) => response.json())
      .then((data) => {
        updateNavbar(data.LoggedIn);
        updateWelcomeSection(data.LoggedIn);
        populatePosts(data.LoggedIn, data.posts);
        populateCategories(data.categories);
        updateUserSection(data.LoggedIn, data.Username);
        updateFilterPostsSection(data.LoggedIn);
      });

    function updateNavbar(loggedIn) {
      const navbar = document.getElementById("navbar-links");
      if (loggedIn) {
        navbar.innerHTML = `
                  <li><a href="/"><i aria-hidden="true"></i> Home</a></li>
                  <li><a href="/create"><i aria-hidden="true"></i> Create A Post</a></li>
                  <li><a href="/logout"><i aria-hidden="true"></i> Log Out</a></li>`;
      } else {
        navbar.innerHTML = `
                  <li><a href="/"><i aria-hidden="true"></i> Home</a></li>
                  <li><a href="/login"><i aria-hidden="true"></i> Login Area</a></li>`;
      }
    }

    function updateWelcomeSection(loggedIn) {
      const joinNowButton = document.getElementById("join-now-button");
      if (!loggedIn) {
        joinNowButton.innerHTML = `
                  <a href="/register">
                      <button class="join92">Join Now</button>
                  </a>`;
      }
    }

    function populatePosts(LoggedIn, posts) {
      const main = document.getElementById("main");
      main.innerHTML = posts
        .map(
          (post) => `
              <div class="question-type2033">
                  <div class="row">
                      <div class="right-description893">
                          <h3><a href="detailsPost/${
                            post.PostID
                          }">${post.Title}</a></h3>
                          <p>${post.Content}</p>
                          <hr>
                          <div class="ques-icon-info3293">
                              <span>${post.Username}</span>
                              <span>${post.FormattedDate}</span>
                              <span>${post.CategoryName}</span>
                          </div>
                          <div class="right-section">
                              ${
                                LoggedIn
                                  ? `
                                <button class="button like" onclick="handleLikeDislike('${post.PostID}', 'like', event)">
                                  <span id='${post.PostID}-like' >👍${post.LikeCount}</span>
                              </button>
                              <button class="button like" onclick="handleLikeDislike('${post.PostID}','dislike', event)">
                                  <span id='${post.PostID}-dislike' >👍${post.DisLikeCount}</span>
                              </button>
                              `
                                  : `
                             <button class="button like">
                                  <span id='${post.PostID}-like' >👍${post.LikeCount}</span>
                              </button>
                              <button class="button like">
                                  <span id='${post.PostID}-dislike' >👍${post.DisLikeCount}</span>
                              </button>
                              `
                              }
                          </div>
                      </div>
                      <div class="ques-type302">
                          <a href="detailsPost/${post.PostID}">
                              <button class="comment-button">${post.CommentCount} Comments</button>
                          </a>
                      </div>
                  </div>
              </div>
          `
        )
        .join("");
    }

    function populateCategories(categories) {
      const categoryList = document.getElementById("category-list");
      categoryList.innerHTML = categories
        .map(
          (category) => `
              <label>
                  <input type="radio" value="${category.ID}" name="categorie" />
                  <span class="custom-checkbox">${category.Name}</span>
              </label>
              <br>
          `
        )
        .join("");
    }

    function updateUserSection(loggedIn, username) {
      const userInfo = document.getElementById("user-info");
      if (loggedIn) {
        userInfo.innerHTML = `
                  <div class="login-part2389">
                      <h4>Welcome, ${username}</h4>
                      <a href="/logout"><button type="button" class="userlogin320">Log Out</button></a>
                  </div>`;
      }
    }
    function updateFilterPostsSection(loggedIn) {
      const filterPostsSection = document.getElementById(
        "filter-posts-section"
      );
      if (loggedIn) {
        filterPostsSection.innerHTML = `
                  <div class="categori-part329">
                      <h4>Filter Posts</h4>
                      <ul>
                          <label>
                              <input type="radio" name="filter" value="created" />
                              <span class="custom-checkbox">Created Posts</span>
                          </label>
                          <label>
                              <input type="radio" name="filter" value="liked" />
                              <span class="custom-checkbox">Liked Posts</span>
                          </label>
                          <button class="resetFilter userlogin320" onclick="Resetfilter()">Reset Filter</button>
                      </ul>
                  </div>
              `;
      } else {
        filterPostsSection.innerHTML = "";
      }
    }

    const btnResetFilter = document.querySelector(".resetFilter")
    console.log(btnResetFilter)

if(btnResetFilter){
        btnResetFilter.addEventListener("click", () => {
            console.log("dfgdfg ")
        const filterby = document.querySelector('input[name="filter"]:checked');
        filterby.checked = false
        handleFilterChange()
    })
}
    
  });

  function handleLikeDislike(postID, action, event) {
    const url = `/${action}/${postID}`;

    fetch(url, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error(`Failed to ${action} the post.`);
        }
        return response.json();
      })
      .then((data) => {
        console.log(data);

        updatePostLikeDislikeCount(
          data.id,
          data.likeCount,
          data.dislikeCount
        );
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  }

  function updatePostLikeDislikeCount(postID, likeCount, dislikeCount) {
    const likeSpan = document.querySelector(`#${CSS.escape(postID)}-like`);
    const dislikeSpan = document.querySelector(`#${CSS.escape(postID)}-dislike`);
    if (likeSpan) likeSpan.textContent = `👍${likeCount}`;
    if (dislikeSpan) dislikeSpan.textContent = `👎${dislikeCount}`;
  }