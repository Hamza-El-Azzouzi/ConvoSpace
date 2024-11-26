
const btnResetCategorie = document.querySelector(".resetCategorie")
function Resetfilter(){
    const filterby = document.querySelector('input[name="filter"]:checked');
    filterby.checked = false
    handleFilterChange()
}



btnResetCategorie.addEventListener("click", () => {
    const categories = document.querySelector('input[name="categorie"]:checked');
    categories.checked = false
    handleFilterChange()
})

function debounce(func, wait) {
    let timeout;
    return function (...args) {
        clearTimeout(timeout);
        timeout = setTimeout(() => func.apply(this, args), wait);
    };
}

function handleFilterChange() {
    const filterby = document.querySelector('input[name="filter"]:checked');
    const categorie = document.querySelector('input[name="categorie"]:checked');
    const pagination = currentPage
    const categoryVal = categorie ? categorie.value : "";
    const filterbyVal = filterby ? filterby.value : "";
    const queryParams = new URLSearchParams({
        filterby: filterbyVal,
        categories: categoryVal,
        pagination :pagination
    });

    fetch('/filters?' + queryParams.toString(), {
            method: 'GET'
        })
        .then(response => response.json())
        .then(data => {
            updateData(data.posts,data.LoggedIn);
            const totalPosts = data.posts.length > 0 ? data.posts[0].totalPosts : 0
            updatePaginationControls(totalPosts)
        })
        .catch((error) => {
            console.error('Error:', error);
        });
}
function updatePaginationControls(totalPages) {
    const pageInfo = document.querySelector("#page-info");
    const nextBtn = document.querySelector("#next-btn");
    const prevBtn = document.querySelector("#prev-btn");
  
    if (!pageInfo || !nextBtn || !prevBtn) {
      console.error("Pagination controls not found in the DOM.");
      return;
    }
  
    // Update page info
    pageInfo.textContent = `Page ${currentPage + 1} of ${Math.ceil(
      totalPages / postsPerPage
    )}`;
  
    // Enable/disable Next button
    if (currentPage + 1 >= Math.ceil(totalPages / postsPerPage)) {
      nextBtn.disabled = true;
    } else {
      nextBtn.disabled = false;
    }
  
    // Enable/disable Prev button
    if (currentPage === 0) {
      prevBtn.disabled = true;
    } else {
      prevBtn.disabled = false;
    }
  }

document.body.addEventListener("change", debounce(handleFilterChange, 200));

const updateData = (data,LoggedInP) => {
    const mainDiv = document.getElementById("main");
    mainDiv.innerHTML = "";

    if (!data || data.length === 0) {

        const noResultsDiv = document.createElement("div");
        noResultsDiv.classList.add("no-results");
        noResultsDiv.textContent = "No Results Found.";
        mainDiv.appendChild(noResultsDiv);
        return;
    }

    data.forEach(post => {
        const postDiv = document.createElement("div");
        postDiv.classList.add("question-type2033");
    
        const rowDiv = document.createElement("div");
        rowDiv.classList.add("row");
    
        const col1 = document.createElement("div");
        col1.classList.add("col-md-1");
        col1.innerHTML = `
            <div class="left-user12923 left-user12923-repeat">
                <a href="#"><i class="fa fa-check" aria-hidden="true"></i></a>
            </div>`;
        const col9 = document.createElement("div");
        col9.classList.add("col-md-9");
        col9.innerHTML = `
            <div class="right-description893">
                <div id="que-hedder2983">
                    <h3><a href='detailsPost/${post.PostID}' target="_blank" id="title">${post.Title}</a></h3>
                </div>
                <div class="ques-details10018">
                    <p id="content">${post.Content}</p>
                </div>
                <hr>
                <div class="ques-icon-info3293">
                    <i class="fa fa-user" aria-hidden="true"> ${post.Username}</i>
                    <i class="fa fa-clock-o" aria-hidden="true"> ${post.FormattedDate}</i>
                    <i class="fa fa-hashtag" aria-hidden="true"> ${post.CategoryName}</i>
                </div>
                <div class="right-section">
                    ${LoggedInP ? `
                        <button
                      class="button like"
                      onclick="handleLikeDislike('${post.PostID}', 'like', event)"
                    >
                      <span id="${post.PostID}-like"
                        >👍${post.LikeCount}</span
                      >
                    </button>
                    <button
                      class="button like"
                      onclick="handleLikeDislike('${post.PostID}','dislike', event)"
                    >
                      <span id="${post.PostID}-dislike"
                        >👎${post.DisLikeCount}</span
                      >
                    </button>
                    ` : `
                        <button class="button like">
                            <span id="${post.PostID}-like">👍${post.LikeCount}</span>
                        </button>
                        <button class="button like">
                            <span id="${post.PostID}-dislike">👎${post.DisLikeCount}</span>
                        </button>
                    `}
                </div>
            </div>`;
    
        const col2 = document.createElement("div");
        col2.classList.add("col-md-2");
        col2.innerHTML = `
            <div class="ques-type302">
                 <a href="detailsPost/${post.PostID}">
                    <button class="comment-button">${post.CommentCount} Comments</button>
                </a>
            </div>`;
    
        rowDiv.appendChild(col1);
        rowDiv.appendChild(col9);
        rowDiv.appendChild(col2);
        postDiv.appendChild(rowDiv);
        mainDiv.appendChild(postDiv);
       
    });
    main.innerHTML += `
    <div class="pagination">
        <button id="prev-btn" class="button" onclick="Previous()">Previous</button>
        <span id="page-info"></span>
        <button id="next-btn" class="button next-btn" onclick="Next()">Next</button>
      </div>`;
    
};
function NextFilter() {
    currentPage++;
    console.log(currentPage)
    handleFilterChange()
    scrollToTop()
  }
  function PreviousFilter() {
    if (currentPage > 1) {
      currentPage--;
      console.log(currentPage)
      handleFilterChange()
      scrollToTop()
    }
  }