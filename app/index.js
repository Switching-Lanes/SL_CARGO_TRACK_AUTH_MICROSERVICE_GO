import React, { useState } from "react";
import axios from "axios";
import styles from "./loginpag.css";
import { Navigate } from "react-router-dom";

const LoginPage = () => {
  const [login, setLogin] = useState("");
  const [password, setPassword] = useState("");
  const [registrationSuccess, setRegistrationSuccess] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();

    const data = {
      login: login,
      password: password,
    };
    try {
      const response = await axios.post(
        "https://vm4506017.43ssd.had.wf/api/users/",
        data
      );

      if (response.status >= 200 && response.status < 300) {
        console.log("Регистрация прошла успешно");
        setRegistrationSuccess(true);
      } else {
        console.error("Ошибка регистрации");
      }
    } catch (error) {
      console.error("Ошибка соединения", error);
    }
  };

  if (registrationSuccess) {
    return <Navigate to='/' />;
  }

  const handleLogin = async (e) => {
    e.preventDefault();

    const data = {
      "login": login,
      "password": password,
    };

    try {
      const response = await axios.post(
        "https://vm4506017.43ssd.had.wf/api/token/login/",
      );

      if (response.data.success) {
        console.log("Вход выполнен успешно");
      } else {
        console.error("Ошибка входа");
      }
    } catch (error) {
      console.error("Ошибка соединения", error);

      console.error("Ошибка входа");
    }
  };

  return (
    <div style={{ marginTop: "80px" }}>
      <div className='lmain'>
        <input type='checkbox' id='chk' aria-hidden='true' />
        <div className='llogin'>
          <form onSubmit={handleLogin} className='lform'>
            <label for='chk' aria-hidden='true'>
              Log in
            </label>
            <input
              className='linput'
              name='email'
              placeholder='Имя'
              required=''
              value={login}
              onChange={(e) => setLogin(e.target.value)}
            />
            <input
              className='linput'
              type='password'
              name='pswd'
              placeholder='Пароль'
              required=''
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
            <button type='submit'>Войти</button>
          </form>
        </div>

        <div className='lregister'>
          <form onSubmit={handleSubmit} className='lform'>
            <label for='chk' aria-hidden='true'>
              Регистрации
            </label>
            <input
              className='linput'
              placeholder='Имя'
              value={login}
              autoComplete='username'
              onChange={(e) => setLogin(e.target.value)}
              required=''
            />
            <input
              className='linput'
              placeholder='Пароль'
              required
              value={password}
              autoComplete='current-password'
              onChange={(e) => setPassword(e.target.value)}
            />
            <button>Зарегистрироваться</button>
          </form>
        </div>
      </div>
    </div>
  );
};

export default LoginPage;