import { API } from "../services/API.js";
import { MovieItem } from "./MovieItem.js";


export class HomePage extends HTMLElement {

    constructor() {
        super();
    }


    // ASYNC/AWAIT VERSION (modern way)
    async render() {
        // This will WAIT for the API call to complete
        const topMovies = await API.getTopMovies();
        renderMoviesInList(topMovies, this.querySelector("#top-20 ul"));

        // This will WAIT for the API call to complete
        const randomMovies = await API.getRandomMovies();
        renderMoviesInList(randomMovies, this.querySelector("#random ul"));


        function renderMoviesInList(movies, ul) {
            ul.innerHTML = "";
            movies.forEach(movie => {
                const li = document.createElement("li");
                li.appendChild(new MovieItem(movie));
                ul.appendChild(li);
            });
        }
    }
    // LIFECYCLE METHOD - Called automatically by browser when element is added to DOM
    connectedCallback() {
        const template = document.getElementById("template-home");
        // cloneNode(true) creates a deep copy of the template's content, including all child elements
        const content = template.content.cloneNode(true);
        this.appendChild(content);  // Add the HTML structure

        this.render();  // Start loading data
    }
}

customElements.define("home-page", HomePage);