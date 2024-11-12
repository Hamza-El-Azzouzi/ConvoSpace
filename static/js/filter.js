const btnResetFilter = document.querySelector(".resetFilter")
const btnResetCategorie = document.querySelector(".resetCategorie")

if(btnResetFilter){
        btnResetFilter.addEventListener("click", () => {
        const filterby = document.querySelector('input[name="filter"]:checked');
        filterby.checked = false
        handleFilterChange()

    })
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
    console.log("was clicked here")
    const categoryVal = categorie ? categorie.value : "";
    const filterbyVal = filterby ? filterby.value : "";
    const queryParams = new URLSearchParams({
        filterby: filterbyVal,
        categories: categoryVal
    });

    console.log(queryParams.toString());

    fetch('/filters?' + queryParams.toString(), {
            method: 'GET'
        })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            updateData(data);
        })
        .catch((error) => {
            console.error('Error:', error);
        });
}

document.body.addEventListener("change", debounce(handleFilterChange, 200));

const updateData = (data) => {
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
                    <h3><a href='detailsPost/${post.post_id}' target="_blank" id="title">${post.title}</a></h3>
                </div>
                <div class="ques-details10018">
                    <p id="content">${post.content}</p>
                </div>
                <hr>
                <div class="ques-icon-info3293">
                    <i class="fa fa-user" aria-hidden="true"> ${post.username}</i>
                    <i class="fa fa-clock-o" aria-hidden="true"> ${post.formatted_date}</i>
                    <i class="fa fa-hashtag" aria-hidden="true"> ${post.category_names}</i>
                </div>
            </div>`;

        const col2 = document.createElement("div");
        col2.classList.add("col-md-2");
        col2.innerHTML = `
            <div class="ques-type302">
                <a href='detailsPost/${post.post_id}'>
                    <button type="button" class="q-type238">
                        <i class="fa fa-comment" aria-hidden="true"> ${post.comment_count} Comments</i>
                    </button>
                </a>
            </div>`;

        rowDiv.appendChild(col1);
        rowDiv.appendChild(col9);
        rowDiv.appendChild(col2);
        postDiv.appendChild(rowDiv);
        mainDiv.appendChild(postDiv);
    });
};