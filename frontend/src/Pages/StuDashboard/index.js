import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios'; // Import Axios
import "./index.css";
function StuDashboard() {
//  const [user, setUser] = useState({});
 const [courses, setCourses] = useState([]);
 const navigate = useNavigate();

 useEffect(() => {
    const userId = sessionStorage.getItem("userId");
    // if(userId){
    //     const fetchUserInfo = async() => {
    //         try {
    //             const response = await axios.get(`/api/users/${userId}`); 
    //             setUser(response.data);
    //         } catch (error) {
    //             console.error("Error fetching user info:", error);
    //         }
    //     };

    //     fetchUserInfo();
    // }
    // Use Axios to fetch courses from the backend
    const fetchCourses = async () => {
      try {
        // Adjusted to fetch courses specific to the user
        // const response = await axios.get(`/api/users/${userId}/courses`);
        const response = await axios.post(`http://localhost:5000/api/student/listcourses`, {"studentid": userId});
        setCourses(response.data); // Assuming the data is directly in response.data
      } catch (error) {
        console.error('Error fetching courses:', error);
      }
    };

    fetchCourses();
 }, []);

 return (
    <div className="course-list-container">
      {/* <div>
      <h2>Welcome, {user.name}</h2>
      </div> */}
      <button onClick={() => navigate('/buycourse')}>Buy Course</button>
      {courses.length > 0 ? (
        courses.map((course) => (
        <div key={course.ID} className="course-item">
          <h3>{course.Coursename.String}</h3>
          <button onClick={() => {navigate("/flashcards"); sessionStorage.setItem("courseid", course.ID);}}>Go to FlashCards</button>
          <button disabled>Join Lecture</button>
        </div>
      ))
      ): (
        <p>You dont have any courses! Buy them!</p>
      )}
    </div>
 );
}

export default StuDashboard;
