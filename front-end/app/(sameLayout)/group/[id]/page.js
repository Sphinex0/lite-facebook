"use client"
import styles from './group.module.css'
import { useState, useEffect } from 'react';
import { use } from "react";
export default function ShowGroup({ params }) {
  const id = use(params).id;

  const [groupData, setGroupData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  console.log(JSON.stringify({ id: parseInt(id) }))
  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('http://localhost:8080/api/group', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          credentials: 'include',
          body: JSON.stringify({ id: parseInt(id) }),
        });
        JSON.stringify({ id })
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }

        const data = await response.json();
        setGroupData(data);
      } catch (error) {
        setError(error.message);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, [id]);

  if (loading) return <h1>Loading...</h1>;
  if (error) return <h1>Error: {error}</h1>;

  return (
    <div>
    <div className={styles.profileHeader}>
      <div className={styles.basicInfo}>
      <div className={styles.p10}>
        <img className={styles.image}></img>
      </div>
      <div className={styles.g2}>
        
        <h1 className={styles.title}>{groupData.title}</h1>
        <span className={styles.nickname}>{groupData.description}</span><br/>
        <span className={styles.followText}>{new Date(groupData.created_at * 1000).toLocaleDateString()}</span><br/>
      </div>
      <div className={`${styles.g1} ${styles.btnSection}`}>
        <button className={styles.editProfileBtn}></button>
      </div>
      </div>
      <div className={styles.profileNav}>
        <ul className={styles.navUl}>
          <li className={`${styles.navLi} ${styles.active}`}>Posts</li>
          <li className={styles.navLi}>events</li>
          <li className={styles.navLi}>mombers</li>
        </ul>

      </div>
    </div>

  </div>


    // <div>
    //   <h1>Group ID: {id}</h1>
    //   {groupData ? (
    //     <div>
    //       <h2>Group Information:</h2>
    //       <h1 >{groupData.title}</h1>
    //       <span >{groupData.description}</span><br />
    //       <span >{new Date(groupData.created_at * 1000).toLocaleDateString()}</span>
    //       {/* <span >281 followers</span><br /> */}
    //     </div>
    //   ) : (
    //     <p>No group data found.</p>
    //   )}
    // </div>
  );

}
