import { User } from "../index.js";
import { Post } from "../post/post.js";

/**
 * manage the display of the post
 * 
 * @param {boolean} filtered 
 */
export const HomeCenter = async (filtered = false) => {
    let data
    
    if (filtered) {
        data = await fetch('http://localhost:8080/getcategories05842365165', { method: 'POST', body: JSON.stringify([User.Id, User.Username, User.Age, User.Gender, User.FirstName, User.LastName, User.Email, User.Password]) }).then(async data => data.json())
    } else {
        data = await fetch('http://localhost:8080/getallposts05842365165', { method: 'POST', body: JSON.stringify([User.Id, User.Username, User.Age, User.Gender, User.FirstName, User.LastName, User.Email, User.Password]) }).then(async data => data.json())
    }
    document.getElementsByClassName('bodyCenter')[0].innerHTML = ''

    data.Posts.forEach(post => {
        const previousPost = document.createElement('div')
        previousPost.setAttribute('class', 'previousPost')
        previousPost.addEventListener('click', () => {
            Post(post.Id)
        })
        document.getElementsByClassName('bodyCenter')[0].appendChild(previousPost)

        let categories = post.NameCategories.reduce((acc, current) => acc + ', ' + current, "")
        categories = categories.slice(2)


        const firstline = document.createElement('div')
        firstline.setAttribute('class', 'firstline')
        previousPost.appendChild(firstline)

        const img = document.createElement('img')
        img.src = post.CreatorImageProfile
        firstline.appendChild(img)

        const info = document.createElement('p')
        info.setAttribute('class', 'info')
        info.textContent = `${post.NameCreator} - ${post.CreationDate}`
        firstline.appendChild(info)

        const likes_dislikes = document.createElement('p')
        likes_dislikes.setAttribute('class', 'like_dislikes')
        likes_dislikes.innerHTML = `&#x1F44D;${post.Likes.Quantity}   &#x1F44E; ${post.Dislikes.Quantity}`
        firstline.appendChild(likes_dislikes)

        const nameCategorie = document.createElement('p')
        nameCategorie.setAttribute('class', 'nameCategorie')
        nameCategorie.innerHTML = `
        <i>Categories:</i>
        ${categories}
        `
        previousPost.appendChild(nameCategorie)

        const br = document.createElement('br')
        previousPost.appendChild(br)

        const namePost = document.createElement('p')
        namePost.setAttribute('class', 'namePost')
        namePost.innerHTML = `
            <i>Name:</i>
            ${post.Name}
        `
        previousPost.appendChild(namePost)

        const descriptionPostSpan = document.createElement('span')
        descriptionPostSpan.innerHTML = '<i>Description:</i>'
        previousPost.appendChild(descriptionPostSpan)

        const descriptionPostPre = document.createElement('pre')
        descriptionPostPre.setAttribute('class', 'descriptionPost')
        descriptionPostPre.innerHTML = post.Description
        previousPost.appendChild(descriptionPostPre)
    });
}
