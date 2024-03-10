  import React from 'react';
  import { useNavigate } from 'react-router-dom';
  import './index.css';

  function Landing() {
  const navigate = useNavigate();

  return (
      <div className="landing-container"> 
        <h2>Lingua Connect</h2>
        <button onClick={() => navigate('/studentLogin')}>Login as Student</button>
        <button onClick={() => navigate('/studentRegister')}>Regsiter as Student</button>
        <button onClick={() => navigate('/tutorLogin')}>Login as Tutor</button>
        <button onClick={() => navigate('/tutorRegister')}>Register as Tutor</button>
      </div>
  );
  }

  export default Landing;
