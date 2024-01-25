// MovieCard.js

import React from "react";

const MovieCard = ({ movie, onClick }) => {
  return (
    <div className="movie-card" onClick={onClick}>
      <h3>{movie.title}</h3>
      <p>Movie ID: {movie.ID}</p>
      <p>Release Date: {new Date(movie.release_date).toLocaleDateString()}</p>
    </div>
  );
};

export default MovieCard;
