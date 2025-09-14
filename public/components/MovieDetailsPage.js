import { API } from "../services/API.js";
import AnimatedLoading from "./AnimatedLoading.js";

export default class MovieDetailsPage extends HTMLElement {
    id = null;
    movie = null;

    async render() {
        try {
            this.movie = await API.getMovieById(this.id);
        } catch {
            alert("Error fetching movie details");
            return;
        }
        const template = document.getElementById("template-movie-details");

        if (!template) {
            console.error("Movie details template not found!");
            return;
        }

        const content = template.content.cloneNode(true);
        this.appendChild(content);

        // Small delay to ensure DOM is fully updated
        setTimeout(() => {
            this.querySelector("h2").textContent = this.movie.title;
            this.querySelector("h3").textContent = this.movie.tagline;
            this.querySelector("img").src = this.movie.poster_url;
            this.querySelector("#trailer").dataset.url = this.movie.trailer_url;
            this.querySelector("#overview").textContent = this.movie.overview;
            this.querySelector("#metadata").innerHTML = `
                <dt>Release Date</dt>
                <dd>${this.movie.release_year}</dd>
                <dt>Score</dt>
                <dd>${this.movie.score} / 10</dd>
                <dt>Popularity</dt>
                <dd>${this.movie.popularity}</dd>
            `;

            const ulGenres = this.querySelector("#genres");
            ulGenres.innerHTML = "";
            this.movie.genre.forEach(genre => {
                const li = document.createElement("li");
                li.textContent = genre.name;
                ulGenres.appendChild(li);
            });

            const ulCast = this.querySelector("#cast");

            if (!ulCast) {
                console.error("Cast element (#cast) not found in the DOM!");
                return;
            }

            ulCast.innerHTML = "";

            if (!this.movie.casting || this.movie.casting.length === 0) {
                console.warn("No cast data available for movie:", this.movie?.title);
                const li = document.createElement("li");
                li.textContent = "No cast information available";
                li.style.color = "#888";
                li.style.fontStyle = "italic";
                ulCast.appendChild(li);
                return;
            }

            this.movie.casting.forEach(actor => {
                const li = document.createElement("li");
                li.innerHTML = `
                    <img src="${actor.image_url ?? '/images/generic_actor.jpg'}" alt="Picture of ${actor.last_name}">
                    <p>${actor.first_name} ${actor.last_name}</p>
                `;
                ulCast.appendChild(li);
            });
        }, 10); // Small delay
    }

    connectedCallback() {
        // Get movie ID from URL or attribute
        this.id = this.params[0];
        this.render();
    }
}
customElements.define("movie-details-page", MovieDetailsPage);