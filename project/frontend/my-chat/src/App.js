import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import WelcomePage from "./components/WelcomePage";
import CreateUsername from "./components/CreateUsername";



function App() {
  return (
    
   
      
    <BrowserRouter>
      <Routes>
        <Route exact path="/" element={< WelcomePage/>}> </Route>
        <Route path="/createUsername" element={< CreateUsername/>}></Route>  
        
        
        
      </Routes> 
    </BrowserRouter>
    
  )
}



export default App;