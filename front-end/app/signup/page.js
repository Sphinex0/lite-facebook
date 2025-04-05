'use client'

import { useState, useEffect } from 'react'
import { useRouter } from 'next/navigation'
import Link from 'next/link'
import styles from "./signup.module.css"
import { FetchApi } from '../helpers'
import FakeNavbar from '../(sameLayout)/_components/navbar/FakeNavbar'

export default function SignupPage() {

  const [form, setForm] = useState({
    email: '',
    password: '',
    firstName: '',
    lastName: '',
    dob: '',
    avatar: null,
    nickname: '',
    aboutMe: ''
  })
  const redirect = useRouter()

  const [error, seterror] = useState('')

  const router = useRouter()

  const handleChange = e => {
    setForm({ ...form, [e.target.name]: e.target.name === "dob" ? (new Date(e.target.value)).getTime() : e.target.value })
  }

  const handleFileChange = e => {
    setForm({ ...form, avatar: e.target.files[0] })
  }

  const handleSignup = async e => {
    e.preventDefault() 
    const data = new FormData()
    for (let i in form) {
      data.append(i, form[i])
    }

    try {
      const response = await FetchApi('/api/signup', redirect, {
        method: 'POST',
        body: data,
      })

      if (response.status == 200) {
        router.push('/login')
      } else {
        const data = await response.json()
        seterror(data)
      }
    } catch (error) {
      seterror("network accured error!")
    }
  }
  /*
    useEffect(() => {
      if (error) {
        const timer = setTimeout(() => seterror(''), 3000)
        return () => clearTimeout(timer)
      }
    }, [error])
  */
  return (
    <>
      <FakeNavbar />
      <form onSubmit={handleSignup}>
        <div className={styles.container}>
          <div className={styles.formBox}>
            <h2 className={styles.heading}>Sign Up</h2>

            {/* Error Popup */}
            {error && <div className={styles.errorPopup}>{error}</div>}
            {error && console.log(error)}
            {
              ['email', 'password', 'firstName', 'lastName', 'dob'].map(field => (
                <div key={field} className={styles.inputGroup}>
                  <input
                    type={
                      field === 'dob'
                        ? 'date'
                        : field === 'password'
                          ? 'password'
                          : 'text'
                    }
                    name={field}
                    placeholder={field.charAt(0).toUpperCase() + field.slice(1)}
                    onChange={handleChange}
                    className={styles.inputField}
                  />
                </div>
              ))}

            {/* Optional Fields */}
            <div className={styles.inputGroup}>
              <label className='file-label'>Upload Avatar (Optional)</label>
              <input
                type='file'
                name='avatar'
                className={`${styles.inputField} ${styles.fileInput}`}
                onChange={handleFileChange}
              />
            </div>

            <div className={styles.inputGroup}>
              <input
                type='text'
                name='nickname'
                placeholder='Nickname (Optional)'
                onChange={handleChange}
                className={styles.inputField}
              />
            </div>

            <textarea
              name='aboutMe'
              placeholder='About Me (Optional)'
              className={`${styles.inputField} ${styles.textarea}`}
              onChange={handleChange}
            ></textarea>

            <button type='submit' className={styles.submitBtn}>
              Sign Up
            </button>
            <div className='login-link'>
              <p>
                Don't have an account? <Link href='/login'>Login in here</Link>
              </p>
            </div>
          </div>
        </div>
      </form>
    </>
  )
}
