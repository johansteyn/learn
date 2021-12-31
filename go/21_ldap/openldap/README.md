# Test LDAP server

These scripts can be used to test Ops Manager - LDAP integrations. It will spin up a Docker container which contains a pre-configured instance of OpenLDAP.

This docker container is a configuration of [docker-openldap](https://github.com/osixia/docker-openldap).

## Example test accounts (global owners)

```
admin:cukes#123!
cnelson:Password1!
bbloom:Password1!
cmintz:Password1!
mms-global-owner@acme.com:abc
```

You can find the full user list in [ldif/2-users.ldif](ldif/2-users.ldif) and the available group in [ldif/3-groups.ldif](ldif/3-groups.ldif).

You can use the following command to generate password hash for any new users you may create:

```
perl -MMIME::Base64 -MDigest::MD5 -e '$ctx = Digest::MD5->new; $ctx->add("replacemewithyourpasswordgs"); print "{MD5}" . encode_base64($ctx->digest)'
```

### Quickstart

1\. Install [Docker](https://www.docker.com/get-docker)

2\. Run `make fresh` to build and run the image

3\. [optional] Execute the the LDAP directory integration tests in MMS:

```
bazel test //server/src/test/com/xgen/svc/mms/svc/user:UserSvcLdapDirectoryIntTests --test_output=streamed --test_arg=--jvm_flag=-DuseLocalLdapServer=true
```

4\. Start MMS with LDAP configured to use the local environment

```
bazel run //server:om -- --jvm_flags="-Dmms.extraPropFile=conf-local-ldap-environment.properties"
```

### Running with LDAPS (SSL/TLS)

To run in TLS, you will need the correct certs, signed by a CA. `ego` has a function, `generate_ssl_certs` that will do this for you.

Otherwise, you will need a self signed certificate ("rootCA"), and a certificate signed by that CA. The signed certificate must have a CN that matches `hostname -f` of the machine running docker.

These commands will set these certs up for you, if run from this directory:

Create CA cert:
```
openssl genrsa 2048 > certs/host.key
openssl req -new -x509 -nodes -sha256 -days 365 -key host.key -out certs/host.cert
```

Generate ldap server certs from CA:
```
openssl genrsa -out certs/openldap-test-server.key 2048
openssl req -new -sha256 -key certs/openldap-test-server.key -subj "/C=US/ST=NY/O=MyOrg, Inc./CN=$(hostname -f)" -out certs/openldap-test-server.csr
openssl x509 -req -in certs/openldap-test-server.csr -CA certs/host.cert -CAkey certs/host.key -CAcreateserial -out certs/openldap-test-server.crt -days 500 -sha256
```

These certs will then need to be added to your machines trusted certificate store, or specified directly in the mms configuration.

To test the certs are trusted on your machine, run `openssl s_client -connect localhost:636`. The return code will be 0 if the certificate is trusted.

See `ego`'s `_configure_ldap_user_service` for example mms settings that will work with this setup.

### Additional commands

- Test the OpenLDAP deployment
  `ldapsearch -h localhost -p 1389 -x -w MmsIsAw3s0me -D cn=admin,dc=com -b dc=com '(uid=admin)' member dn`

- Test memberOf. memberOf is a necessary openldap feature for listing reverse group membership. This command should list all groups that user1 belongs to.
  `ldapsearch -h localhost -p 1389 -b dc=babypearfoo,dc=com -x -w MmsIsAw3s0me -D cn=admin,dc=com '(uid=user1)' memberof dn`

- See the available OpenLDAP build options
  `make`

- Encode a new password for a LDAP user
`perl -MMIME::Base64 -MDigest::MD5 -e '$ctx = Digest::MD5->new; $ctx->add("cukes#123!"); print "{MD5}" . encode_base64($ctx->digest)'`

### References

- https://wiki.corp.mongodb.com/display/MMS/MMS+OpenLDAP+Test+Server
