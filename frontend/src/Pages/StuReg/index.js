import { useEffect, useState } from "react";
import axios from "axios";
import "./index.css";
// import { Link } from "react-router-dom";
import { useNavigate } from "react-router-dom";

const StudentRegister = () => {
 const navigate = useNavigate();
 const [email, setEmail] = useState("");
 const [password, setPassword] = useState("");
 const [showPassword, setShowPassword] = useState(false);
//  const [log, setLog] = useState(0);
 const [errorMessage, setErrorMessage] = useState(""); // State to hold error message

 const handleEmailChange = (e) => {
    setEmail(e.target.value);
 };

 const handlePasswordChange = (e) => {
    setPassword(e.target.value);
 };

 const togglePasswordVisibility = () => {
    setShowPassword(!showPassword);
 };

 const register = async () => {
    let data = { "username": email, "password": password };
    // console.log(data);
    try {
      let req = await axios.post("http://localhost:5000/api/auth/sturegister", data);
      // req = req.data.res;
      console.log(req);
      if (req) {
        const userId = req.data.ID;
        // sessionStorage.setItem("userId", userId);
        // sessionStorage.setItem("isLoggedIn", 1);
        alert("You will be directed to the login page")
        navigate("/");
      } else {
        setErrorMessage("Invalid username or password.");
      }
    } catch (error) {
      console.log("Login error:", error);
      setErrorMessage("An error occurred during login.");
    }
 };

 const handleSubmit = (e) => {
    e.preventDefault();
    register();
 };

//  useEffect(() => {
//     const timer = setTimeout(() => {
//       setLog(0);
//     }, 5000);
//     return () => clearTimeout(timer); // Cleanup function to clear the timer
//  }, []);

 return (
    <div className="login-form-container">
      <p>Student Regsiter Page</p>
      <form onSubmit={handleSubmit}>
        <label>
          Email:
          <input
            type="email"
            value={email}
            onChange={handleEmailChange}
            required
          />
        </label>
        <label>
          Password:
          <input
            type={showPassword ? "text" : "password"}
            value={password}
            onChange={handlePasswordChange}
            required
          />
        </label>
        <button type="button" onClick={togglePasswordVisibility}>
          {showPassword ? "Hide" : "Show"} Password
        </button>
        <button type="submit">Login</button>
      </form>
      {errorMessage && <p className="error-message">{errorMessage}</p>}
    </div>
 );
};

export default StudentRegister;