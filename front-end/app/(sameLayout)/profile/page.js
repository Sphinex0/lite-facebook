import styles from './profile.module.css'
const Users= () =>{
  return (
    <div>
      <div className={styles.profileHeader}>
        <div className={styles.basicInfo}>
        <div className={styles.p10}>
          <img className={styles.image}></img>
        </div>
        <div className={styles.g2}>
          
          <h1 className={styles.title}>ADNANE EL MIR</h1>
          <span className={styles.nickname}>@Sphinex</span><br/>
          <span className={styles.followText}>281 followers</span><br/>
          <span className={styles.followText}>81 following</span>
        </div>
        <div className={`${styles.g1} ${styles.btnSection}`}>
          <button className={styles.editProfileBtn}>edit Profile</button>
        </div>
        </div>
        <div className={styles.profileNav}>
          <ul className={styles.navUl}>
            <li className={`${styles.navLi} ${styles.active}`}>Posts</li>
            <li className={styles.navLi}>About</li>
            <li className={styles.navLi}>Following</li>
            <li className={styles.navLi}>Followers</li>
          </ul>

        </div>
      </div>

    </div>
  )
}
export default Users;