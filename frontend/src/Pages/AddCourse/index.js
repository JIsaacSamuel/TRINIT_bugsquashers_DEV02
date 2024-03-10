import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import "./index.css";
function AddCourse() {
  const navigate = useNavigate();
 const [language, setLanguage] = useState('');
 const [price, setPrice] = useState('');
 const [tutorId, setTutorId] = useState(sessionStorage.getItem('userId'));
 const [courseName, setCourseName] = useState('');

 const handleLanguageChange = (e) => {
    setLanguage(e.target.value);
 };

 const handlePriceChange = (e) => {
    setPrice(e.target.value);
 };

 const handleCourseNameChange = (e) => {
    setCourseName(e.target.value);
 };

 const handleSubmit = async (e) => {
    e.preventDefault();

    const course = {
      "ID": tutorId,
      "lang" : language,
      "price" : parseInt(price,10),
      "coursename" : courseName,
    };
    console.log(course);
    try {
      const response = await axios.post('http://localhost:5000/api/tutor/createcourse', course);
      console.log(response.data);
      alert("Courses added")
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
      <h2>Add New Course</h2>
      <form onSubmit={handleSubmit}>
        <label>
          Language:
          <input type="text" value={language} onChange={handleLanguageChange} required />
        </label>
        <label>
          Price:
          <input type="number" value={price} onChange={handlePriceChange} required />
        </label>
        <label>
          Tutor ID:
          <input type="text" value={tutorId} onChange={setTutorId} disabled />
        </label>
        <label>
          Course Name:
          <input type="text" value={courseName} onChange={handleCourseNameChange} required />
        </label>
        <button type="submit">Add Course</button>
      </form>
      
    </div>
 );
}

export default AddCourse;
