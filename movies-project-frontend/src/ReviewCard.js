// ReviewCard.js

import React from "react";

const reviewCardStyle = {
  border: "1px solid #ccc",
  borderRadius: "8px",
  padding: "10px",
  margin: "10px 0",
  maxWidth: "400px",
  textAlign: "left",
};

const strongStyle = {
  marginRight: "5px",
};

const ReviewCard = ({ review }) => {
  return (
    <div style={reviewCardStyle}>
      <strong style={strongStyle}>ID:</strong> {review.ID},{" "}
      <strong style={strongStyle}>Movie ID:</strong> {review.movieId},{" "}
      <strong style={strongStyle}>User ID:</strong> {review.userId},{" "}
      <strong style={strongStyle}>Rating:</strong> {review.rating},{" "}
      <strong style={strongStyle}>Text:</strong> {review.reviewText}
    </div>
  );
};

export default ReviewCard;
