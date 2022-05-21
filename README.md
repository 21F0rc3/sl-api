<div id="top"></div>

<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![License][license-shiled]][license-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">

  <a href="https://github.com/21F0rc3/sl-api/">
    <img src="https://user-images.githubusercontent.com/57480698/168663263-6b838ee1-e0f9-442d-bf8e-d0f7bb79d87b.png" alt="Logo" width="150" height="150">
  </a>

  <h3 align="center">Smart Lion API</h3>

  <p align="center">
    A GO API to integrate an IoT system on the disposal of used kitchen oil.
    <br />
    <a href="https://github.com/21F0rc3/sl-api/"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/21F0rc3/sl-api/">SmartLion API</a>
    ·
    <a href="https://github.com/fredmnpinto/smart_lion_mobile/">SmartLion Mobile</a>
    ·
    <a href="https://github.com/21F0rc3/sl-api/issues">Report Bug</a>
    ·
    <a href="https://github.com/21F0rc3/sl-api/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#authors">Authors</a>
    </li>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
    </li>
  </ol>
</details>


### Authors:
- Gabriel Fernandes (40344@ufp.edu.pt)
- Frederico Pinto (39822@ufp.edu.pt)
  
## About The Project

Neste  projecto  os  alunos  irão  ter  oportunidade  de  trabalhar  com  várias  tecnologias, nomeadamente, sistemas inteligentes que combinam a utilização de tecnologias Internet of Things (IoT), classificadores usando Machine Learning (ML) e desenvolvimento de uma API de Backend (python flask ou go lang) e respectivo Frontend (app móvel - Flutter/Kotliin). 

O projeto envolve o desenvolvimento de um sistema inteligente para automatizar a recolha de Óleos Alimentares Usados (OAU) numa rede de oleões espalhados pelo país. No projeto irão ser utilizados 3 depósitos de recolha de OAU que já estão apetrechados com um micro-controlador ESP32 e vários sensores (cf. temperatura, fluxo, nível, turbidez) para recolher algumas das propriedades dos líquidos depositados. A informação recolhida é processada localmente e enviada através de um protocolo IoT para um servidor de Cloud (API backend). Os dados recolhidos serão posteriormente utilizados para treinar algoritmos de machine learning com o intuito de permitir distinguir futuros depósitos de líquidos usando a mesma arquitectura.

O sistema deverã ainda possuir uma app móvel que permite identificar univocamente os utilizadores que fazem os depósitos (e.g. usando BLE ou QR-codes) e que recebem pontos de fidelização por cada depósito válido (i.e. depósito de OAU). Para aferir a qualidade de cada depósito deverá utilizar-se os classificadores previamente treinados, que localmente irão determinar se o depósito é valido (cf. OAU) ou se não é válido (cf. contem percentagens significativas de água ou óleo de carro usado misturados).

O  sistema  deverá  ainda  possuir  uma  app  móvel  que  permite  identificar univocamente os utilizadores que fazem os depósitos (e.g. usando BLE ou QR-codes) e que recebem pontos de fidelização por cada depósitos válido (i.e. depósito de OAU). Para aferir  a  qualidade  de  cada  depósito  deverá  utilizar-se  os  classificadores  previamente treinados, que localmente irão determinar se o depósito é válido (cf. OAU) ou se não é válido (cf. contem percentagens significativas de água ou óleo de carro usado misturados 

## Getting Started


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/21F0rc3/sl-api?style=for-the-badge
[contributors-url]: https://github.com/21F0rc3/sl-api/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/21F0rc3/sl-api?style=for-the-badge
[forks-url]: https://github.com/21F0rc3/sl-api//network/members
[stars-shield]: https://img.shields.io/github/stars/21F0rc3/sl-api?style=for-the-badge
[stars-url]: https://github.com/21F0rc3/sl-api/stargazers
[issues-shield]: https://img.shields.io/github/issues/21F0rc3/sl-api?style=for-the-badge
[issues-url]: https://github.com/21F0rc3/sl-api/issues
[license-shield]: https://img.shields.io/github/license/21F0rc3/sl-api
[license-url]: https://github.com/21F0rc3/sl-api/blob/master/LICENSE
