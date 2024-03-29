FROM owasp/modsecurity:v2-ubuntu-apache
MAINTAINER Jing Lu lvjing5@huawei.com

ARG COMMIT=v3.1/dev
ARG REPO=SpiderLabs/owasp-modsecurity-crs
ENV PARANOIA=1

RUN a2enmod rewrite

RUN apt-get update && \
    apt-get -y install python git ca-certificates iproute2 vim

RUN cd /opt && \
  git clone https://github.com/${REPO}.git owasp-modsecurity-crs-3.1 && \
  cd owasp-modsecurity-crs-3.1 && \
  git checkout -qf ${COMMIT}

RUN cd /opt && \
  cp -R /opt/owasp-modsecurity-crs-3.1/ /etc/apache2/modsecurity.d/owasp-crs/ && \
  mv /etc/apache2/modsecurity.d/owasp-crs/crs-setup.conf.example /etc/apache2/modsecurity.d/owasp-crs/crs-setup.conf && \
  cd /etc/apache2/modsecurity.d && \
  printf "include modsecurity.d/owasp-crs/crs-setup.conf\ninclude modsecurity.d/owasp-crs/rules/*.conf" > include.conf && \
  sed -i -e 's/SecRuleEngine DetectionOnly/SecRuleEngine On/g' /etc/apache2/modsecurity.d/modsecurity.conf && \
  a2enmod proxy proxy_http

COPY proxy.conf           /etc/apache2/modsecurity.d/proxy.conf
COPY docker-entrypoint.sh /

RUN chmod 777 /docker-entrypoint.sh

COPY .htaccess            /var/www/html/.htaccess
COPY apache2.conf         /etc/apache2/apache2.conf

EXPOSE 80

ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["apachectl", "-D", "FOREGROUND"]
