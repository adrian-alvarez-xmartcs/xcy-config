import { BrowserRouter, Route, Routes } from 'react-router-dom'
import './App.css'
import LoginPage from './pages/Login'
import { UserProvider } from './contexts/UserContext'

function App() {

  return (
    <>
    <UserProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<LoginPage />}/>
        </Routes>
      </BrowserRouter>
    </UserProvider>
    </>
  )
}

export default App
