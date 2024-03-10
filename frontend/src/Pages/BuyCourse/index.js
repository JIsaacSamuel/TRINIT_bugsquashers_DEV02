// BuyCourse.js
import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import "./index.css";
function BuyCourse() {
  const navigate = useNavigate();
//  const [searchTerm, setSearchTerm] = useState('');
 const [budget, setBudget] = useState('');
 const [language, setLanguage] = useState('');
 const [courses, setCourses] = useState([]);

 const handleSearch = async () => {
    // Implement the search logic here
    // For example, make an API call to your backend with the searchTerm, budget, and language
    // Then, update the courses state with the response
    try {
        const response = await axios.post('http://localhost:5000/api/student/requiredcourse', {
          // searchTerm,
          "lang":language,
          "price": budget,
        });
        setCourses(response.data); // Assuming the data is directly in response.data
     } catch (error) {
        console.error('Error fetching courses:', error);
     }
 };
 const handleBuy = async (courseId) => {
  const userId = sessionStorage.getItem("userId");

  if (!userId) {
     console.error("User ID not found in session storage.");
     alert("Course Bought");
     return;
  }
 
  try {
     const response = await axios.post('http://localhost:5000/api/student/coursesubsribing', {
       "studentid": userId,
       "courseid": courseId
     });
     console.log(response.data);
     // Handle success, e.g., show a success message or update the UI
    //  navigate('/stuDashboard');
  } catch (error) {
     console.error('Error buying course:', error);
     // Handle error, e.g., show an error message
  }
 };
 

 return (
    <div className="course-search-form-container">
    <button className="back-button" onClick={() => navigate('/stuDashboard')}>Back to Dashboard</button>
      <h2>Type in the Details of the Course</h2>
      {/* <input
        type="text"
        placeholder="Search for a tutor"
        value={searchTerm}
        onChange={(e) => setSearchTerm(e.target.value)}
      /> */}
      <input
        type="number"
        placeholder="Budget"
        value={budget}
        onChange={(e) => setBudget(e.target.value)}
      />
      <select value={language} onChange={(e) => setLanguage(e.target.value)}>
        <option value="">Select Language</option>
        <option value="Tamil">Tamil</option>
        <option value="Spanish">Spanish</option>
        <option value="French">French</option>
        {/* Add more languages as needed */}
      </select>
      <button onClick={handleSearch}>Search</button>
      {courses.length > 0 ? (
 courses.map((course) => (
    <div key={course.ID}>
      <h3>{course.name}</h3>
      <p>Tutor: {course.Name.String}</p>
      <p>Course: {course.Coursename.String}</p>
      <p>Price: {course.Price}</p>
      <button onClick={() => handleBuy(course.ID)}>Buy</button>
    </div>
  ))
) : (
 <p>No courses for the above criterion.</p>
)}

    </div>
 );
}

export default BuyCourse;
