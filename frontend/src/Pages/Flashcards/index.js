import React, { useState, useEffect } from 'react';
// import { useParams } from 'react-router-dom';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import "./index.css";
function Flashcards() {
  const navigate = useNavigate();
  //  const { courseId } = useParams();
 const [flashcards, setFlashcards] = useState([]);
 const courseId = sessionStorage.getItem("courseid");
 useEffect(() => {
    const fetchFlashcards = async () => {
      try {
        const response = await axios.post("/api", courseId);
        setFlashcards(response.data);
      } catch (error) {
        console.error('Error fetching flashcards:', error);
      }
    };

    fetchFlashcards();
 }, [courseId]);

 return (
    <div>
      <button className="back-button" onClick={() => navigate('/stuDashboard')}>Back to Dashboard</button>
      <h2>Flashcards for Course {courseId}</h2>
      {flashcards.map((flashcard, index) => (
        <div key={index} className="flashcard">
          <div className="flashcard-front">
            {flashcard.frontText}
          </div>
          <div className="flashcard-back">
            {flashcard.backText}
          </div>
        </div>
      ))}
    </div>
 );
}

export default Flashcards;
