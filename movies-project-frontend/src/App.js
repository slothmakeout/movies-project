import React, { useState, useEffect } from "react";
import axios from "axios";
import MovieCard from "./MovieCard";
import ReviewCard from "./ReviewCard";

const App = () => {
  const [movies, setMovies] = useState([]);
  const [selectedMovieId, setSelectedMovieId] = useState(null);
  const [selectedMovie, setSelectedMovie] = useState(null);
  const [selectedReviewId, setSelectedReviewId] = useState(null);
  const [selectedReview, setSelectedReview] = useState(null);
  const [reviews, setReviews] = useState([]);
  const [searchedReviews, setSearchedReviews] = useState([]);
  const [newReview, setNewReview] = useState({
    movieID: null,
    rating: null,
    reviewText: "",
  });

  useEffect(() => {
    axios
      .get("http://localhost:8080/movies")
      .then((response) => setMovies(response.data))
      .catch((error) => console.error("Error fetching movies:", error));

    axios
      .get("http://localhost:8080/reviews")
      .then((response) => setReviews(response.data))
      .catch((error) => console.error("Error fetching reviews:", error));
  }, []);

  const handleSearchById = () => {
    if (selectedMovieId) {
      axios
        .get(`http://localhost:8080/movies/${selectedMovieId}`)
        .then((response) => setSelectedMovie(response.data))
        .catch((error) => {
          console.error("Error fetching movie by ID:", error);
          setSelectedMovie(null);
        });

      setSelectedReviewId(null);
      setSelectedReview(null);
      setSearchedReviews([]);
    }
  };

  const handleSearchReviewById = () => {
    if (selectedReviewId) {
      axios
        .get(`http://localhost:8080/reviews/${selectedReviewId}`)
        .then((response) => setSelectedReview(response.data))
        .catch((error) => {
          console.error("Error fetching review by ID:", error);
          setSelectedReview(null);
        });
    }
  };

  const handleSearchReviewsByMovieId = () => {
    if (selectedMovieId) {
      axios
        .get(`http://localhost:8080/movies/${selectedMovieId}/reviews`)
        .then((response) => {
          setSearchedReviews(response.data);

          axios
            .get("http://localhost:8080/reviews")
            .then((response) => setReviews(response.data))
            .catch((error) => console.error("Error fetching reviews:", error));
        })
        .catch((error) => {
          console.error("Error fetching reviews by movie ID:", error);
          setSearchedReviews([]);
        });

      setSelectedReviewId(null);
      setSelectedReview(null);
    }
  };

  const handlePostReview = () => {
    if (newReview.movieID && newReview.rating && newReview.reviewText) {
      axios
        .post("http://localhost:8080/reviews", {
          movieID: newReview.movieID,
          rating: newReview.rating,
          reviewText: newReview.reviewText,
        })
        .then((response) => {
          axios
            .get("http://localhost:8080/reviews")
            .then((response) => setReviews(response.data))
            .catch((error) => console.error("Error fetching reviews:", error));

          setNewReview({
            movieID: null,
            rating: null,
            reviewText: "",
          });
        })
        .catch((error) => console.error("Error posting review:", error));
    }
  };

  const handleUpdateReview = () => {
    if (selectedReviewId && newReview.rating !== null) {
      axios
        .put(`http://localhost:8080/reviews/${selectedReviewId}`, {
          rating: newReview.rating,
          reviewText: newReview.reviewText,
        })
        .then((response) => {
          axios
            .get("http://localhost:8080/reviews")
            .then((response) => setReviews(response.data))
            .catch((error) => console.error("Error fetching reviews:", error));

          setNewReview({
            movieID: null,
            rating: null,
            reviewText: "",
          });

          setSelectedReviewId(null);
          setSelectedReview(null);
        })
        .catch((error) => console.error("Error updating review:", error));
    }
  };

  const handleDeleteReview = () => {
    if (selectedReviewId) {
      axios
        .delete(`http://localhost:8080/reviews/${selectedReviewId}`)
        .then((response) => {
          axios
            .get("http://localhost:8080/reviews")
            .then((response) => setReviews(response.data))
            .catch((error) => console.error("Error fetching reviews:", error));

          setNewReview({
            movieID: null,
            rating: null,
            reviewText: "",
          });

          setSelectedReviewId(null);
          setSelectedReview(null);
        })
        .catch((error) => console.error("Error deleting review:", error));
    }
  };

  return (
    <div className="app">
      <h1>Movies</h1>

      <div className="movie-list">
        {movies.map((movie) => (
          <MovieCard
            key={movie.ID}
            movie={movie}
            onClick={() => setSelectedMovieId(movie.ID)}
          />
        ))}
      </div>

      <div className="search-form">
        <label>Search Movie by ID:</label>
        <input
          type="number"
          value={selectedMovieId || ""}
          onChange={(e) => setSelectedMovieId(e.target.value)}
        />
        <button onClick={handleSearchById}>Search</button>
      </div>

      {selectedMovie && (
        <MovieCard key={selectedMovie.ID} movie={selectedMovie} />
      )}

      <h2>All Reviews</h2>
      <div className="reviews-list">
        <ul>
          {reviews.map((review) => (
            <ReviewCard key={review.ID} review={review} />
          ))}
        </ul>
      </div>

      <div className="search-form">
        <label>Search Review by ID:</label>
        <input
          type="number"
          value={selectedReviewId || ""}
          onChange={(e) => setSelectedReviewId(e.target.value)}
        />
        <button onClick={handleSearchReviewById}>Search</button>
      </div>

      {selectedReview && (
        <ReviewCard key={selectedReview.ID} review={selectedReview} />
      )}

      <div className="search-form">
        <label>Search Reviews by Movie ID:</label>
        <input
          type="number"
          value={selectedMovieId || ""}
          onChange={(e) => setSelectedMovieId(e.target.value)}
        />
        <button onClick={handleSearchReviewsByMovieId}>Search</button>
      </div>

      {searchedReviews.length > 0 && (
        <div className="reviews-container">
          <h2>All Reviews</h2>
          <div className="reviews-list">
            <ul>
              {searchedReviews.map((review) => (
                <ReviewCard key={review.ID} review={review} />
              ))}
            </ul>
          </div>
        </div>
      )}

      <div className="add-review-form">
        <h2>Add a Review</h2>
        <label>Movie ID:</label>
        <input
          type="number"
          value={newReview.movieID || ""}
          onChange={(e) =>
            setNewReview({ ...newReview, movieID: parseInt(e.target.value) })
          }
        />
        <label>Rating:</label>
        <input
          type="number"
          step="0.1"
          value={newReview.rating || ""}
          onChange={(e) =>
            setNewReview({ ...newReview, rating: parseFloat(e.target.value) })
          }
        />
        <label>Review Text:</label>
        <textarea
          value={newReview.reviewText || ""}
          onChange={(e) =>
            setNewReview({ ...newReview, reviewText: e.target.value })
          }
        ></textarea>
        <button onClick={handlePostReview}>Add Review</button>
      </div>

      <div className="update-review-form">
        <h2>Update Review</h2>
        <label>Review ID:</label>
        <input
          type="number"
          value={selectedReviewId || ""}
          onChange={(e) => setSelectedReviewId(e.target.value)}
        />
        <label>New Rating:</label>
        <input
          type="number"
          step="0.1"
          value={newReview.rating || ""}
          onChange={(e) =>
            setNewReview({ ...newReview, rating: parseFloat(e.target.value) })
          }
        />
        <label>New Text:</label>
        <textarea
          value={newReview.reviewText || ""}
          onChange={(e) =>
            setNewReview({ ...newReview, reviewText: e.target.value })
          }
        ></textarea>
        <button onClick={handleUpdateReview}>Update Review</button>
      </div>

      <div className="delete-review-form">
        <h2>Delete Review</h2>
        <label>Review ID:</label>
        <input
          type="number"
          value={selectedReviewId || ""}
          onChange={(e) => setSelectedReviewId(e.target.value)}
        />
        <button onClick={handleDeleteReview}>Delete Review</button>
      </div>
    </div>
  );
};

export default App;
