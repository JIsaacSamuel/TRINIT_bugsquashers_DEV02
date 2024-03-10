import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import "./index.css";
function CreateFlashCards() {
    const courseId = sessionStorage.getItem("courseid")
  const navigate = useNavigate();
//  const [language, setLanguage] = useState('');
  const [word1, setword1] = useState('');

  const [mean1, setmean1] = useState('');
  
 
 const handleWord1 = (e) => {
    setword1(e.target.value);
 };

 const handleMean1 = (e) => {
    setmean1(e.target.value);
 };



 const handleSubmit = async (e) => {
    e.preventDefault();

    const card = {
      "ID": courseId,
      "text": word1,
      "meaning": mean1,

    };
    console.log(card);
    try {
      const response = await axios.post('http://localhost:5000/api/', card);
      console.log(response.data);
      alert("FlashCard added")
      // Navigate('/stuDashboard');
      // Handle success
    } catch (error) {
      console.error('Error adding course:', error);
      // Handle error
    }
 };

 return (
    <div className="add-course-form-container">
      <button className="back-button" onClick={() => navigate('/tutorDashboard')}> Back to dashboard</button>
      <h2>Add New Flashcard</h2>
      <form onSubmit={handleSubmit}>
        <label>
          Text 1:
          <input type="text" value={word1} onChange={handleWord1} required />
        </label>
        <label>
          Meaning 1:
          <input type="text" value={mean1} onChange={handleMean1} required />
        </label>
        
        <button type="submit">Add Flashcard</button>
      </form>
      
    </div>
 );
}

export default CreateFlashCards;
