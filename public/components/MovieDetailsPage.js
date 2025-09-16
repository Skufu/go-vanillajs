import { API } from "../services/API.js";
import AnimatedLoading from "./AnimatedLoading.js";

export default class MovieDetailsPage extends HTMLElement {
    id = null;
    movie = null;

    async render() {
        try {
            this.movie = await API.getMovieById(this.id);
        } catch (error) {
            console.error("Error fetching movie details:", error);
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
            try {
                this.querySelector("h2").textContent = this.movie.title || "Unknown Title";
                this.querySelector("h3").textContent = this.movie.tagline || "";
                this.querySelector("img").src = this.movie.poster_url || "/images/generic_actor.jpg";
                this.querySelector("#trailer").dataset.url = this.movie.trailer_url || "";
                this.querySelector("#overview").textContent = this.movie.overview || "No overview available";

                this.querySelector("#metadata").innerHTML = `
                    <dt>Release Date</dt>
                    <dd>${this.movie.release_year || "Unknown"}</dd>
                    <dt>Score</dt>
                    <dd>${this.movie.score || "N/A"} / 10</dd>
                    <dt>Popularity</dt>
                    <dd>${this.movie.popularity || "N/A"}</dd>
                `;

                const ulGenres = this.querySelector("#genres");
                ulGenres.innerHTML = "";
                if (this.movie.genre && Array.isArray(this.movie.genre)) {
                    this.movie.genre.forEach(genre => {
                        const li = document.createElement("li");
                        li.textContent = genre?.name || "Unknown Genre";
                        ulGenres.appendChild(li);
                    });
                }

                const ulCast = this.querySelector("#cast");

                if (!ulCast) {
                    console.error("Cast element (#cast) not found in the DOM!");
                    return;
                }

                ulCast.innerHTML = "";

                if (!this.movie.casting || !Array.isArray(this.movie.casting) || this.movie.casting.length === 0) {
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
                        <img src="${actor?.image_url ?? '/images/generic_actor.jpg'}" alt="Picture of ${actor?.last_name || 'Unknown'}">
                        <p>${actor?.first_name || ''} ${actor?.last_name || 'Unknown'}</p>
                    `;
                    ulCast.appendChild(li);
                });
            } catch (error) {
                console.error("Error rendering movie details:", error);
            }
        }, 10); // Small delay
    }

    connectedCallback() {
        // Get movie ID from URL or attribute
        if (!this.params || !this.params[0]) {
            console.error("Movie ID not found in params:", this.params);
            return;
        }
        this.id = this.params[0];
        this.render();
    }
}
customElements.define("movie-details-page", MovieDetailsPage);