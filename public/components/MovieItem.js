class MovieItem extends HTMLElement {
    constructor(movie) {
        super();
        this.movie = movie;
    }
    connectedCallback() {
        const url = "/movies/" + this.movie.id;
        this.innerHTML = `
            <article onclick="app.Router.go('${url}')" style="cursor: pointer;">
                <img src="${this.movie.poster_url}" alt="${this.movie.title} Poster">
                <p>${this.movie.title} (${this.movie.release_year})</p>
            </article>
        `

    }
}


customElements.define("movie-item", MovieItem);

export default MovieItem;