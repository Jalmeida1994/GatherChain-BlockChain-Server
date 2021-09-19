
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Thanks again! Now go create something AMAZING! :D
-->



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
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/Jalmeida1994/GatherChain-BlockChain-Server">
    <img src="images/Logo-02.png" alt="Logo" width="120">
  </a>

  <h3 align="center">GatherChain BlockChain Server</h3>

  <p align="center">
    BlockChain server for the GatherChain solution.
    <br />
    <a href="https://github.com/Jalmeida1994/GatherChain-BlockChain-Server/blob/master/README.md"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/Jalmeida1994/GatherChain-BlockChain-Server/issues">Report Bug</a>
    ·
    <a href="https://github.com/Jalmeida1994/GatherChain-BlockChain-Server/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

GatherChain is the solution created for my Master Thesis: __Tracing Responsibility in Evolution of Model's Life Cycle in Collaborative Projects in Education__ in the Informatics' Department of __NOVA School of Science and Technology__ made with professor __Vasco Amaral__'s guidance.
In the paper, it is proposed a blockchain-based solution for version control of model-driven engineering artefacts.  The goal is to facilitate collaboration in a multi-user area, like the education field, and track changes in a trusted and secure manner. This solution is based on using the Hyperledger Fabric Network to govern and regulate file version control functions among students and teachers.
This repository is a part of the larger GatherChain solution.

The other GatherChain projects are:
* __GatherChain Desktop Client__: https://github.com/Jalmeida1994/GatherChain-DesktopClient
* __GatherChain Web Server__: https://github.com/Jalmeida1994/GatherChain-Web-Server
* __GatherChain Admin Commands__: https://github.com/Jalmeida1994/GatherChain-Admin-Commands
* __GatherChain ARM Template__: https://github.com/Jalmeida1994/GatherChain-ARM-Template

This repository serves as the code for the blockchain network to be installed in the server.

### Built With

The network is built with:
* [Hyperledger-Fabric](https://www.hyperledger.org/use/fabric)
    * [Minifabric](https://github.com/hyperledger-labs/minifabric)
* [Shell Scripts](https://www.shellscript.sh)
    * [GNU Bash](https://www.gnu.org/software/bash/)
* [Go](https://golang.org)


<!-- GETTING STARTED -->
## Getting Started

The blockchain network is used automatically by the [GatherChain ARM Template](https://github.com/Jalmeida1994/GatherChain-ARM-Template) and abstracted by the [GatherChain Web Server](https://github.com/Jalmeida1994/GatherChain-Web-Server). If you want to deploy this network in your environment of choice then this section it'll be shown how to get started. Only tested in `ubuntu 20.04`. 

### Prerequisites

* [Bash shell](https://www.gnu.org/software/bash/)
* [Go](https://golang.org/doc/install)
* [Docker](https://docs.docker.com/get-docker/)
* [net-tools](https://zoomadmin.com/HowToInstall/UbuntuPackage/net-tools)
* [jq](https://stedolan.github.io/jq/download/)
* [ca-certificates](https://www.techrepublic.com/article/how-to-install-ca-certificates-in-ubuntu-server/)
* [software-properties-common](https://zoomadmin.com/HowToInstall/UbuntuPackage/software-properties-common)


### Installation


1. Clone the repo
   ```
   git clone https://github.com/Jalmeida1994/GatherChain-BlockChain-Server.git
   ```
   

<!-- USAGE EXAMPLES -->
## Usage

The blockchain can be communicated via [GatherChain Web Server](https://github.com/Jalmeida1994/GatherChain-Web-Server) using the [GatherChain Admin Commands](https://github.com/Jalmeida1994/GatherChain-Admin-Commands) and [GatherChain Desktop App](https://github.com/Jalmeida1994/GatherChain-DesktopClient).

If you want to use the blockchain network inside the server, you can follow the next commands:


<!-- USAGE EXAMPLES -->
### Initialize the network
    ```
    ./commands/init.sh ${author} ${groupName} ${commitHash}
    ```
### Create a channel in the network
    ```
    ./commands/createchannel.sh ${author} ${groupName} ${commitHash}
    ```
### Push changes to the blockchain network
    ```
    ./commands/push.sh ${author} ${groupName} ${commitHash}
    ```
### Clears the blockchain network
    ```
    ./commands/clear.sh
    ```
### Get history of transactions (Prints commits and authors to console)
    ```
    ./commands/gethistory.sh ${groupName}
    ```

To check the commands used by the scripts and create your own check the [commands/fabric_commands.md](https://github.com/Jalmeida1994/GatherChain-BlockChain-Server/blob/master/commands/fabric_commands.md)

<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/Jalmeida1994/GatherChain-BlockChain-Server/issues) for a list of proposed features (and known issues).



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact

João Almeida - [@João Almeida](https://www.linkedin.com/in/jo%C3%A3o-almeida-525476125/) - jcfd.almeida@campus.fct.unl.pt

Project Link: [https://github.com/Jalmeida1994/GatherChain-BlockChain-Server](https://github.com/Jalmeida1994/GatherChain-BlockChain-Server)



<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements
* [FCT-UNL](https://www.fct.unl.pt/)
* [Professor Vasco Amaral](https://docentes.fct.unl.pt/vma/)
* [Tong Li](https://github.com/litong01/)


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/Jalmeida1994/GatherChain-BlockChain-Server.svg?style=for-the-badge
[contributors-url]: https://github.com/Jalmeida1994/GatherChain-BlockChain-Server/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/Jalmeida1994/GatherChain-BlockChain-Server.svg?style=for-the-badge
[forks-url]: https://github.com/Jalmeida1994/GatherChain-BlockChain-Server/network/members
[stars-shield]: https://img.shields.io/github/stars/Jalmeida1994/GatherChain-BlockChain-Server.svg?style=for-the-badge
[stars-url]: https://github.com/Jalmeida1994/GatherChain-BlockChain-Server/stargazers
[issues-shield]: https://img.shields.io/github/issues/Jalmeida1994/GatherChain-BlockChain-Server.svg?style=for-the-badge
[issues-url]: https://github.com/Jalmeida1994/GatherChain-BlockChain-Server/issues
[license-shield]: https://img.shields.io/github/license/Jalmeida1994/GatherChain-BlockChain-Server.svg?style=for-the-badge
[license-url]: https://github.com/Jalmeida1994/GatherChain-BlockChain-Server/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/jo%C3%A3o-almeida-525476125/
