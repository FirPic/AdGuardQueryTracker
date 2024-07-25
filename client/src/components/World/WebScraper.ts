/* eslint-disable no-console */
/* eslint-disable @typescript-eslint/no-unused-vars */
import axios from 'axios';


function isValidIPv4OrDomain(entree) {
    // Expression régulière pour une adresse IPv4
    const regexIPv4 = /^(\d{1,3}\.){3}\d{1,3}$/;
    // Expression régulière pour un nom de domaine
    const regexDomaine = /^[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;

    // Enlever le port et les paramètres
    const entreeSansPortParams = entree.replace(/(:\d+)?(\/.*)?$/, '');

    // Vérifier si c'est une IPv4 valide
    if (regexIPv4.test(entreeSansPortParams)) {
        return entreeSansPortParams;
    }
    // Vérifier si c'est un domaine valide
    if (regexDomaine.test(entreeSansPortParams)) {
        return entreeSansPortParams;
    }
    return null;
}

export async function getInfoByIP(request) {
    if (!isValidIPv4OrDomain(request)) {
        return { error: 'Invalid IP or domain', url: request };
    }
    try {
        const response = await axios({
            method: 'get',
            url: `https://www.infobyip.com/ip-${request}.html`,
            responseType: 'text',
            headers: {
                'Access-Control-Allow-Origin': '*',
                'User-Agent':
                    'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36',
                Accept: 'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8',
                'Accept-Language': 'en-US,en;q=0.9',
                'Accept-Encoding': 'gzip, deflate, br',
                Connection: 'keep-alive',
                'Upgrade-Insecure-Requests': '1',
            },
        });

        console.log(`scraper content : ${response.data}`);

        return {
            ip: request,
            domain: '', 
            continent: '', 
            'continent-code': '', 
            country: '', 
            'country-code': '', 
            isp: '', 
        };
    } catch (error) {
        console.error(error);
        return { error: 'Failed to fetch data' };
    }
}
