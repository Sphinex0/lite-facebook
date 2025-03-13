'use client';

import styles from './groupe.module.css';
import { useState } from 'react';

const Group = () => {
  const [compContent, setCompContent] = useState(null); // Store data directly
  const [loading, setLoading] = useState(false); // For loading state

  // Function to handle the link click (fetch data)
  const handleLinkClick = async (content) => {
    setLoading(true); // Show loading spinner or text
    try {
      const response = await fetch(`http://localhost:8080/api/` + content, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include',
      });

      if (!response.ok) {
        throw new Error('Network response was not ok');
      }

      const data = await response.json();
      console.log(data);

      setCompContent(data); 
    } catch (error) {
      console.error('Fetch error:', error);
      setCompContent(null); 
    } finally {
      setLoading(false);
    }
  };

  const renderTable = (data) => {
    if (!data || data.length === 0) {
      return <p>No data available</p>;
    }

    return (
      <table className={styles.table}>
        <thead>
          <tr>
            <th>Title</th>
            <th>Description</th>
          
          
        
          </tr>
        </thead>
        <tbody>
          {data.map((row, index) => (
            <tr key={index}>

              <td>{row.title}</td>
              <td>{row.description}</td>

             
            </tr>
          ))}
        </tbody>
      </table>
    );
  };

  return (
    <div>
      <h1>Groups</h1>
      <div className={styles.typegroupes}>
        <a href="#!" onClick={() => handleLinkClick('groups')}>
          <div>All Groups</div>
        </a>
        <a href="#!" onClick={() => handleLinkClick('members')}>
          <div>Your Groups</div>
        </a>
      </div>

      <div className={styles.comp}>
        {loading && <p>Loading...</p>}
        {compContent && renderTable(compContent)}
      </div>
    </div>
  );
};

export default Group;
