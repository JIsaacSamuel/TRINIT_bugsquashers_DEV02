import React, { useState, useEffect } from 'react';
// import { useParams } from 'react-router-dom';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import "./index.css";
function CourseDetails() {
//  const { courseId } = useParams();
 const [students, setStudents] = useState([]);
 const navigate = useNavigate();

 useEffect(() => {
    const fetchStudents = async () => {
      const courseId = sessionStorage.getItem("courseid");
      try {
        const response = await axios.post('http://localhost:5000/api/tutor/getstudents', {
          "courseid": courseId,
        });
        setStudents(response.data);
      } catch (error) {
        console.error('Error fetching students:', error);
      }
    };

    fetchStudents();
 }, []);

 return (
    <div className="course-details-container">
     <button className="back-button" onClick={() => navigate('/tutorDashboard')}>Back to Dashboard</button>
      <h2>Course Details</h2>
      <button onClick={() => {navigate('/createFlashcards')}}>Add Flashcards</button>
      {students.length > 0 ? (
        students.map((student) => (
        <div key={student.ID} className="student-item">
          <p>{student.Name.String}</p>
          <button disabled>Join the Meet</button>
        </div>
      ))
      ):(
        <p>There are no students enrolled for this course</p>
      )}
     
    </div>
 );
}

export default CourseDetails;
