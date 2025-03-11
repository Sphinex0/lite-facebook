export default function MainLayout({ children }) {
  /**
   *         <main>
        
        {children}
      </main>
   */
  return (
    <main>
      <div className="container">
        <div className="left">
          <p>left side</p>
        </div>

        <div className="middle">
          {children}
        </div>

        <div className="right">
        <p>right side</p>
        </div>
      </div>
    </main>
  );
}