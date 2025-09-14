import { API } from "../services/API.js";
import MovieItem from "./MovieItem.js";

export default class MoviePage extends HTMLElement {

    async render() {
        try {
            // Get all movies using search with empty query
            const allMovies = await API.searchMovies("", "title", "");

            this.renderMoviesList(allMovies);
        } catch (error) {
            console.error("Error loading movies:", error);
            this.querySelector("#movies-list").innerHTML = "<li>Error loading movies</li>";
        }
    }

    renderMoviesList(movies) {
        const ul = this.querySelector("#movies-list");
        ul.innerHTML = "";

        if (!movies || movies.length === 0) {
            ul.innerHTML = "<li>No movies found</li>";
            return;
        }

        movies.forEach(movie => {
            const li = document.createElement("li");
            li.appendChild(new MovieItem(movie));
            ul.appendChild(li);
        });
    }

    connectedCallback() {
        const template = document.getElementById("template-movies");
        const content = template.content.cloneNode(true);
        this.appendChild(content);

        this.render();
    }
}

customElements.define("movie-page", MoviePage);