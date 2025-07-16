#!/bin/sh
python manage.py migrate --noinput
python manage.py collectstatic --noinput
exec "$@"  # launches gunicorn after setup