import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import "./index.css";

function TutorDashboard() {
 const navigate = useNavigate();
 const [courses, setCourses] = useState([]);
//  const [user, setUser] = useState({});

 useEffect(() => {
    const userId = sessionStorage.getItem("userId");

    // if(userId){
    //     const fetchUserInfo = async() => {
    //         try {
    //             const response = await axios.get(`/api/users/${userId}`);
    //             setUser(response.data);
    //             // Assuming the response data contains both user and courses information
    //             setCourses(response.data.courses);
    //           } catch (error) {
    //             console.error("Error fetching user info:", error);
    //           }
    //     };
    //     fetchUserInfo();
    // }
    const fetchCourses = async () => {
      try{
        const response = await axios.post('http://localhost:5000/api/tutor/listcourse',{"tutorid":userId});
        console.log(response);
        setCourses(response.data);
        // sessionStorage.setItem("userId", userId);
      }catch(error) {
        console.error('error fetching courses: ', error);
      }
    };
    fetchCourses();
 }, []);

 return (
    <div className="course-list-container">
      {/* <h2>Welcome! {user.name}</h2> */}
      <button onClick={() => navigate('/addCourse')}>Add New Courses</button>
      {courses.length > 0 ? (
        courses.map((course) => (
          <div key={course.ID} className="course-item">
            <h3>{course.Coursename.String}</h3>
            <button onClick={() => {navigate("/course"); sessionStorage.setItem("courseid", course.ID);}}>Go to Course Details</button>
          </div>
        ))
      ) : (
        <p>You are not teaching any courses! add them now!</p>
      )}
    </div>
 );
}

export default TutorDashboard;
