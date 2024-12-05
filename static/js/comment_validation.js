// get the content of the comment and sendt it to the backend (handler /sendcomment) with a fetch and post methode
//receive the response as json that contain all the comments of the post and call update function to append them in the html
function SubmitComment(event) {

    event.preventDefault();

    const path = window.location.pathname;
    const currentPostId = path.substring(path.lastIndexOf('/')+1);
    const textarea = document.querySelector('.comment-textarea');

    if (textarea.value.trim() === "") {
        alert('please enter text !!!')
        return
    }
    const currentValue = textarea.value.trim()

    async function fetchData() {

        try {
            const response = await fetch('/sendcomment',{
                method: "POST",
                headers: {
                    'content-Type': 'application/json',
                },
                body: JSON.stringify({content: currentValue, postId: currentPostId })
            })
            if (!response.ok) {
                throw new Error('http error')
            }
            const data = await response.json()
            return data
            textarea.value = ''
        }
    } catch (error) {
        console.error('There was a problem in fetch :', error);
    }
}


//get the comment section div make it empty then loop over comments and append them to the comment section div
// function UpdateComment(comments) {
    
// }
