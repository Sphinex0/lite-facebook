import Profile from "./_leftSide/profile";
import SideBar from "./_leftSide/sideBar";
import "./globals.css";

export default function MainLayout({ children }) {
  return (
    <main>
      <div className="container">
        <div className="left">
          <Profile/>
          <SideBar/>
          <label className="btn btn-primary">Create Post</label>
        </div>
        
        <div className="middle">
          {children}
        </div>
        {/* 
        <div className="right">
        <p>right side</p>
        </div> */}
      </div>
    </main>
  );
}