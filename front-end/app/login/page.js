'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'
import styles from "./login.module.css"
import { FetchApi } from '../helpers'
import Link from 'next/link' 

export default function Login () {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const router = useRouter()

  const [error, seterror] = useState('')

  const handleSubmit = async e => {
    e.preventDefault()

    try {
      const response = await FetchApi('/api/login', router,{
        method: 'POST',
        body: JSON.stringify({ email, password })
      })

      if (response.status == 200) {
        const data = await response.json()
        localStorage.setItem('user', JSON.stringify(data))
        router.push('/')
      } else {
        const data = await response.json()
        seterror(data)
      }
    } catch (error) {
      seterror(error)
    }
  }

  return (
    <div className={styles.container}>
      <div className={styles.formBox}>
        <h2 className={styles.heading}>Login</h2>

        {/* Error Popup */}
        {error && <div className={styles.errorPopup}>{error}</div>}
        <form onSubmit={handleSubmit}>
          {/* Email Input */}
          <div className={styles.inputGroup}>
            <input
              type='text'
              value={email}
              placeholder='User or Email'
              required
              id='username'
              onChange={e => setEmail(e.target.value)}
              className={styles.inputField}
            />
          </div>

          {/* Password Input */}
          <div className={styles.inputGroup}>
            <input
              type='password'
              value={password}
              placeholder='Password'
              required
              id='password'
              onChange={e => setPassword(e.target.value)}
              className={styles.inputField}
            />
          </div>

          {/* Submit Button */}
          <button type='submit' className={styles.submitBtn}>
            Login
          </button>
        </form>

        <div className="signup-link">
          <p>Don't have an account? <Link href="/signup">Sign up here</Link></p>
        </div>
      </div>
    </div>
  )
}