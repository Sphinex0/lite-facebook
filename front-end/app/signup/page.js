'use client'

import { useState , useEffect} from 'react'
import { useRouter } from 'next/navigation'
import './signup.css'

export default function SignupPage () {
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

  const [error, seterror] = useState("")

  const router = useRouter()

  const handleChange = e => {
    setForm({ ...form, [e.target.name]: e.target.value })
  }

  const handleFileChange = e => {
    setForm({ ...form, avatar: e.target.files[0] })
  }

  const handleSignup = async e => {
    e.preventDefault() // Prevent the default form submit behavior
    const data = new FormData()
    for (let i in form) {
      data.append(i, form[i])
    }

    try {
      const response = await fetch('http://localhost:8080/api/signup', {
        method: 'POST',
        body: data,// Send form data as a JSON string
        credentials:"include"
      })

      if (response.status == 200) {
        // If the response is ok, navigate to the homepage
        const data = await response.json()
        sessionStorage.setItem('first_name', data.first_name)
        sessionStorage.setItem('last_name', data.last_name)
        sessionStorage.setItem('Nickname', data.nickname)
        sessionStorage.setItem('Image', data.image)
        router.push('/')
      } else {
        const data = await response.json()
        console.log(data);
                
        seterror(data)
      }
    } catch (error) {
      console.log(data);

      seterror(data)
    }
  }

  useEffect(() => {
    if (error) {
      const timer = setTimeout(() => seterror(""), 3000);
      return () => clearTimeout(timer);
    }
  }, [error]);

  return (
    <form onSubmit={handleSignup}>
      <div className='container'>
        <div className='form-box'>
          <h2>Sign Up</h2>

          {/* Error Popup */}
          {error && <div className='error-popup'>{error}</div>}

          {['email', 'password', 'firstName', 'lastName', 'dob'].map(field => (
            <div key={field} className='input-group'>
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
                className='input-field'
              />
            </div>
          ))}

          {/* Optional Fields */}
          <div className='input-group'>
            <label className='file-label'>Upload Avatar (Optional)</label>
            <input
              type='file'
              name='avatar'
              className='input-field file-input'
              onChange={handleFileChange}
            />
          </div>

          <div className='input-group'>
            <input
              type='text'
              name='nickname'
              placeholder='Nickname (Optional)'
              onChange={handleChange}
              className='input-field'
            />
          </div>

          <textarea
            name='aboutMe'
            placeholder='About Me (Optional)'
            className='input-field textarea'
            onChange={handleChange}
          ></textarea>

          <button type='submit' className='submit-btn'>
            Sign Up
          </button>
        </div>
      </div>
    </form>
  )
}
