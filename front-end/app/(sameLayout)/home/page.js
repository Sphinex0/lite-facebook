"use client";

import Link from 'next/link';
import styles from './home.module.css'
import Requests from './Requests';
import { useState } from 'react';

const Home = () => {

  const [screen, setScreen] = useState("");
  const switchRequest = (sideNav)=>{

    console.log("yes")
    setScreen("home")
  }
  return (
    <div className={styles.container}>
      <aside className={styles.sideNav}>
        {screen == "home"? <Requests/>:""}
        <ul>
          
        <Link href="/profile">
          <li className={styles.navItem}>
            <img className={styles.image}/>
            <p className={styles.userName}>ADNANE ELMIR</p>
          </li>
          </Link>
          <li className={styles.navItem} onClick={(e)=>switchRequest(e.target.parentElement.parentElement)}>
            <img className={`${styles.image} ${styles.Request}`}/>
            <p className={styles.userName}>Requests</p>
          </li>
          <li className={styles.navItem}>
            <img className={`${styles.image} ${styles.group}`}/>
            <p className={styles.userName} >Groups</p>
          </li>
          <li className={styles.navItem}>
            <img className={styles.image}/>
            <p className={styles.userName}>Event</p>
          </li>
        </ul>
      </aside>
      <section>
        <div className={styles.post}>
          <div className={styles.postHeader}>
            <Link href="/profile">
            <li className={styles.navItem}>
              <img className={styles.image}/>
              <p className={styles.userName}>ADNANE ELMIR</p>
            </li>
            </Link>
          </div>

          <div className={styles.postBody}>

          </div>

          <div className={styles.postFooter}>

          </div>

        </div>
      </section>
      <aside>

      </aside>
    </div>
  )
}

export default Home;