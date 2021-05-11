# site.pp
node default {
file { '/tmp/custom-file.txt':
    ensure => 'present',
    content => "Hello World",
  }
}
