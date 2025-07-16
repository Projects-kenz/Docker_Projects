#!/bin/sh

echo "▶️ Running migrations..."
python manage.py migrate

echo "🎨 Collecting static files..."
python manage.py collectstatic --noinput

echo "🚀 Starting server..."
exec "$@"
